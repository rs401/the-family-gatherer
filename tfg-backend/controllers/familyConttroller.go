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
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
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
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
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
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
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
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	db.Where(&models.Thread{ForumID: uint(fid)}).Find(&threads)
	return c.JSON(threads)
}

// Get single thread
func GetThread(c *fiber.Ctx) error {
	db := database.DBConn
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	var thread models.Thread
	db.Find(&thread, id)
	if thread.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	return c.JSON(thread)
}

func NewThread(c *fiber.Ctx) error {
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
	thread := new(models.Thread)
	if err := c.BodyParser(&thread); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	// Check forum exists
	fid, err := strconv.Atoi(c.Params("fid"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	var forum models.Forum
	db.Find(&forum, fid)
	if forum.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	thread.ForumID = uint(fid)
	thread.Forum = forum
	thread.UserID = user.ID
	thread.User = *user
	db.Create(&thread)
	return c.JSON(thread)
}

func DeleteThread(c *fiber.Ctx) error {
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
	// Find the thread
	var thread models.Thread
	db.Find(&thread, id)
	if thread.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	// Check user is owner
	if thread.UserID != user.ID {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	db.Delete(&thread)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "ok",
	})
}

func UpdateThread(c *fiber.Ctx) error {
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
	var oldThread models.Thread
	var updThread = new(models.Thread)
	db.First(&oldThread, id)

	if oldThread.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	// Check user is owner
	if oldThread.UserID != user.ID {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	if err := c.BodyParser(updThread); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	if title := strings.TrimSpace(updThread.Title); title != "" {
		oldThread.Title = updThread.Title
	}
	if body := strings.TrimSpace(updThread.Body); body != "" {
		oldThread.Body = updThread.Body
	}

	db.Save(&oldThread)

	return c.JSON(oldThread)
}

// Posts
// Get all posts
func GetPosts(c *fiber.Ctx) error {
	db := database.DBConn
	var posts []models.Post
	tid, err := strconv.Atoi(c.Params("tid"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	db.Where(&models.Post{ThreadID: uint(tid)}).Find(&posts)
	return c.JSON(posts)
}

// Get single post
func GetPost(c *fiber.Ctx) error {
	db := database.DBConn
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	var post models.Post
	db.Find(&post, id)
	if post.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	return c.JSON(post)
}

func NewPost(c *fiber.Ctx) error {
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
	// Check tid valid
	tid, err := strconv.Atoi(c.Params("tid"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	var thread models.Thread
	db.Find(&thread, tid)
	if thread.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	post := new(models.Post)
	if err := c.BodyParser(&post); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	// Set post ThreadID and UserID
	post.ThreadID = thread.ID
	post.UserID = user.ID
	db.Create(&post)
	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
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
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	// Check post exists
	var post models.Post
	db.Find(&post, id)
	if post.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	// Check user is owner
	if post.UserID != user.ID {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	db.Delete(&post)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "ok",
	})
}

func UpdatePost(c *fiber.Ctx) error {
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
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	var post models.Post
	var updPost = new(models.Post)
	db.First(&post, id)
	if post.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "notfound",
		})
	}
	if err := c.BodyParser(updPost); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "badrequest",
		})
	}
	// Check user is owner
	if post.UserID != user.ID {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	if updPost.Body != "" {
		post.Body = updPost.Body
	}
	db.Save(&post)

	return c.JSON(post)
}
