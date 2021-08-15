package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
)

// Get all posts
func GetPosts(c *fiber.Ctx) error {
	db := database.DBConn
	var posts []models.Post
	db.Find(&posts)
	return c.JSON(posts)
}

// Get single post
func GetPost(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var post models.Post
	db.Find(&post, id)
	if post.Body == "" {
		return c.Status(500).SendString("Post does not exist in the database.")
	}
	return c.JSON(post)
}

func NewPost(c *fiber.Ctx) error {
	db := database.DBConn
	post := new(models.Post)
	if err := c.BodyParser(&post); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&post)
	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var post models.Post
	db.Find(&post, id)
	if post.Body == "" {
		return c.Status(500).SendString("Post does not exist in the database.")
	}

	db.Delete(&post)
	return c.SendString("Post successfully Deleted from database.")
}

func UpdatePost(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var post models.Post
	var updPost = new(models.Post)
	db.First(&post, id)
	if post.Body == "" {
		return c.Status(500).SendString("Post does not exist in the database.")
	}
	if err := c.BodyParser(updPost); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	theId, idErr := strconv.Atoi(id)
	if idErr != nil {
		return c.Status(422).SendString(idErr.Error())
	}
	updPost.ID = uint(theId)
	name := updPost.Body

	if name == "" {
		return c.Status(400).SendString("Fields cannot be empty.")
	}
	db.Save(&updPost)

	return c.JSON(updPost)
}
