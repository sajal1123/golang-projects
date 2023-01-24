package Lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sajal1123/go-fiber/Database"
)

type Lead struct {
	gorm.Model
	Name   string `json:"name"`
	Club   string `json:"club"`
	Number int    `json:"number"`
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := Database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func GetLeads(c *fiber.Ctx) {
	db := Database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func NewLead(c *fiber.Ctx) {
	db := Database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := Database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted.")
}
