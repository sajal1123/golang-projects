package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sajal1123/go-fiber/Database"

	// "github.com/sajal1123/golang-fiber/Database"
	"github.com/sajal1123/go-fiber/Lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", Lead.GetLeads)
	app.Post("/api/v1/lead", Lead.NewLead)
	app.Delete("/api/v1/lead/:id", Lead.DeleteLead)
	app.Get("/api/v1/lead/:id", Lead.GetLead)
}

func initDatabase() {
	var err error
	Database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	Database.DBConn.AutoMigrate(&Lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3335)

	defer Database.DBConn.Close()
}
