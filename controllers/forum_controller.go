package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
)

// Get all forums
func GetForums(c *fiber.Ctx) error {
	db := database.DBConn
	var forums []models.Forum
	db.Find(&forums)
	return c.JSON(forums)
}

// Get single forum
func GetForum(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var forum models.Forum
	db.Find(&forum, id)
	if forum.Name == "" {
		return c.Status(500).SendString("Forum does not exist in the database.")
	}
	return c.JSON(forum)
}

func NewForum(c *fiber.Ctx) error {
	db := database.DBConn
	forum := new(models.Forum)
	if err := c.BodyParser(&forum); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&forum)
	return c.JSON(forum)
}

func DeleteForum(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var forum models.Forum
	db.Find(&forum, id)
	if forum.Name == "" {
		return c.Status(500).SendString("Forum does not exist in the database.")
	}

	db.Delete(&forum)
	return c.SendString("Forum successfully Deleted from database.")
}

func UpdateForum(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var forum models.Forum
	var updForum = new(models.Forum)
	db.First(&forum, id)
	if forum.Name == "" {
		return c.Status(500).SendString("Forum does not exist in the database.")
	}
	if err := c.BodyParser(updForum); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	theId, idErr := strconv.Atoi(id)
	if idErr != nil {
		return c.Status(422).SendString(idErr.Error())
	}
	updForum.ID = uint(theId)
	name := updForum.Name

	if name == "" {
		return c.Status(400).SendString("Fields cannot be empty.")
	}
	db.Save(&updForum)

	return c.JSON(updForum)
}
