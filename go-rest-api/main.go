package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type player struct {
	Id     string `json:"id"`
	Player string `json:"player"`
	Rank   int    `json:"rank"`
}

var players = []player{
	{Id: "3", Player: "Messi", Rank: 1},
	{Id: "1", Player: "Pele", Rank: 2},
	{Id: "2", Player: "Maradona", Rank: 3},
}

func rankGoats(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, players)
}
func addPlayer(context *gin.Context) {
	var newPlayer player

	if err := context.BindJSON(&newPlayer); err != nil {
		return
	}
	players = append(players, newPlayer)
	context.IndentedJSON(http.StatusCreated, newPlayer)
}
func getPlayer(name string) (*player, error) {
	for i, t := range players {
		if t.Player == name {
			return &players[i], nil
		}
	}
	return nil, errors.New("Player not found.")
}

func getInfo(context *gin.Context) {
	name := context.Param("player")
	fmt.Println("\nGot name from context -> ", name, '\n')
	player, err := getPlayer(name)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Player not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, player)
}

func main() {
	// fmt.Println(todos)
	router := gin.Default()
	router.GET("/todos", rankGoats)
	router.GET("/todos/:player", getInfo)
	router.POST("/todos", addPlayer)
	router.Run("localhost:9000")
}
