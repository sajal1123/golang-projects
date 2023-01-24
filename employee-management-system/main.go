package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURL = "mongodb://localhost:27017/" + dbName

type Player struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Jersey   int    `json:"jersey"`
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/player", func(c *fiber.Ctx) error {
		query := bson.D{{}}
		cursor, err := mg.Db.Collection("players").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var players []Player = make([]Player, 0)

		if err := cursor.All(c.Context(), &players); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(players)
	})

	app.Post("/player", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("players")

		player := new(Player)

		if err := c.BodyParser(player); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		player.ID = ""

		insertionRes, err := collection.InsertOne(c.Context(), player)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		filter := bson.D{{Key: "_id", Value: insertionRes.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filter)

		createdPlayer := &Player{}
		createdRecord.Decode(createdPlayer)
		return c.Status(201).JSON(createdPlayer)
	})

	app.Put("/player/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		playerID, err := primitive.ObjectIDFromHex(idParam)

		if err != nil {
			return c.SendStatus(400)
		}
		player := new(Player)
		err = c.BodyParser(player)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: playerID}}
		update := bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{Key: "name", Value: player.Name},
					{Key: "postion", Value: player.Position},
					{Key: "jersey", Value: player.Jersey},
				},
			},
		}
		err = mg.Db.Collection("players").FindOneAndUpdate(c.Context(), query, update).Err()

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(400)
			}
			return c.SendStatus(500)
		}

		player.ID = idParam
		return c.Status(200).JSON(player)
	})

	app.Delete("/player/:id", func(c *fiber.Ctx) error {
		playerID, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}

		query := bson.D{{Key: "_id", Value: playerID}}
		result, err := mg.Db.Collection("players").DeleteOne(c.Context(), &query)
		if err != nil {
			return c.SendStatus(500)
		}

		if result.DeletedCount < 1 {
			return c.SendStatus(404)
		}

		return c.Status(200).JSON("record deleted")
	})

	log.Fatal(app.Listen("localhost:3000"))
}
