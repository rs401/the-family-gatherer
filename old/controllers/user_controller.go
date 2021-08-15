package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
)

// Get all users
func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}

// Get single user
func GetUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user models.User
	db.Find(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User does not exist in the database.")
	}
	return c.JSON(user)
}

func NewUser(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user models.User
	db.Find(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User does not exist in the database.")
	}

	db.Delete(&user)
	return c.SendString("User successfully Deleted from database.")
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user models.User
	var updUser = new(models.User)
	db.First(&user, id)
	if user.DisplayName == "" {
		return c.Status(500).SendString("User does not exist in the database.")
	}
	if err := c.BodyParser(updUser); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	theId, idErr := strconv.Atoi(id)
	if idErr != nil {
		return c.Status(422).SendString(idErr.Error())
	}
	updUser.ID = uint(theId)
	name := updUser.DisplayName

	if name == "" {
		return c.Status(400).SendString("Fields cannot be empty.")
	}
	db.Save(&updUser)

	return c.JSON(updUser)
}
