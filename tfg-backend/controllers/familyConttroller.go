package controllers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/tfg-backend/database"
	"github.com/rs401/TFG/tfg-backend/models"
)

// Forums
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
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	var forum models.Forum
	db.Find(&forum, id)
	if forum.Name == "" {
		return c.Status(500).SendString("Forum does not exist in the database.")
	}
	return c.JSON(forum)
}

func NewForum(c *fiber.Ctx) error {
	cookie := c.Cookies("tfg")
	user := getUserByJwt(cookie)
	if user == nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	db := database.DBConn
	forum := new(models.Forum)
	if err := c.BodyParser(&forum); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	forum.UserID = user.ID
	forum.User = *user
	db.Create(&forum)
	return c.JSON(forum)
}

func DeleteForum(c *fiber.Ctx) error {
	// Get cookie
	cookie := c.Cookies("tfg")
	// Check user
	user := getUserByJwt(cookie)
	if user == nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	// Grab db
	db := database.DBConn
	// Convert string parameter to int
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	// Get the forum
	var forum models.Forum
	res := db.Find(&forum, id)
	if res.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	// Check user is forums owner
	if user.ID != forum.UserID {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "forbidden",
		})
	}
	db.Delete(&forum)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "ok",
	})
}

func UpdateForum(c *fiber.Ctx) error {
	// Get cookie
	cookie := c.Cookies("tfg")
	// Check user
	user := getUserByJwt(cookie)
	if user == nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	// Grab db
	db := database.DBConn
	// Convert string parameter to int
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	// Get original and apply updates
	var forum models.Forum
	var updForum = new(models.Forum)
	res := db.First(&forum, id)
	if res.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	// Parse new values
	if err := c.BodyParser(updForum); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	// Check not empty string
	if updForum.Name == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	// Update forum and save
	forum.Name = updForum.Name
	db.Save(&forum)

	return c.JSON(forum)
}

// Threads
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

// Posts
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
