package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
)

// Get all forums and return with index template
func GetIndex(c *fiber.Ctx) error {
	db := database.DBConn
	var forums []models.Forum
	db.Find(&forums)
	// return c.JSON(forums)
	return c.Render("index", fiber.Map{
		"Title":  "Hello, World!",
		"Forums": forums,
	})
}

// Get About
func GetAbout(c *fiber.Ctx) error {
	db := database.DBConn
	var forums []models.Forum
	db.Find(&forums)
	// return c.JSON(forums)
	return c.Render("about", fiber.Map{
		"Title": "About",
	})
}
