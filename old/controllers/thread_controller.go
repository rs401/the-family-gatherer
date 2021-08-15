package controllers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
)

// Get all threads
func GetThreads(c *fiber.Ctx) error {
	db := database.DBConn
	var threads []models.Thread
	fid, err := strconv.Atoi(c.Params("fid"))
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Where(&models.Thread{ForumID: uint(fid)}).Find(&threads)
	return c.JSON(threads)
}

// Get single thread
func GetThread(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var thread models.Thread
	db.Find(&thread, id)
	if thread.Title == "" {
		return c.Status(500).SendString("Thread does not exist in the database.")
	}
	return c.JSON(thread)
}

func NewThread(c *fiber.Ctx) error {
	db := database.DBConn
	thread := new(models.Thread)
	if err := c.BodyParser(&thread); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	// Check forum exists
	fid, err := strconv.Atoi(c.Params("fid"))
	if err != nil {
		return c.Status(503).SendString(err.Error())
	}
	var forum models.Forum
	db.Find(&forum, fid)
	if forum.Name == "" {
		return c.Status(418).SendString("Forum doesn't exist")
	}
	thread.ForumID = uint(fid)
	thread.Forum = forum
	db.Create(&thread)
	return c.JSON(thread)
}

func DeleteThread(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var thread models.Thread
	db.Find(&thread, id)
	if thread.Title == "" {
		return c.Status(500).SendString("Thread does not exist in the database.")
	}

	db.Delete(&thread)
	return c.SendString("Thread successfully Deleted from database.")
}

func UpdateThread(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var oldThread models.Thread
	var updThread = new(models.Thread)
	db.First(&oldThread, id)
	// Check forum exists
	fid := oldThread.ForumID
	var forum models.Forum
	db.Find(&forum, fid)
	if forum.Name == "" {
		return c.Status(418).SendString("Forum doesn't exist")
	}

	if oldThread.Title == "" {
		return c.Status(500).SendString("Thread does not exist in the database.")
	}
	if err := c.BodyParser(updThread); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	theId, idErr := strconv.Atoi(id)
	if idErr != nil {
		return c.Status(422).SendString(idErr.Error())
	}
	updThread.ID = uint(theId)
	if title := strings.TrimSpace(updThread.Title); title == "" {
		updThread.Title = oldThread.Title
	}
	if body := strings.TrimSpace(updThread.Body); body == "" {
		updThread.Body = oldThread.Body
	}
	if userid := updThread.UserID; userid == 0 {
		updThread.UserID = oldThread.UserID
	}
	if forumid := updThread.ForumID; forumid == 0 {
		updThread.ForumID = oldThread.ForumID
	}

	db.Save(&updThread)

	return c.JSON(updThread)
}
