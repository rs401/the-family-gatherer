package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rs401/TFG/config"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
	"github.com/shareed2k/goth_fiber"
)

// Get all forums and return with index template
func GetIndex(c *fiber.Ctx) error {
	db := database.DBConn
	var forums []models.Forum
	db.Find(&forums)

	session, err := goth_fiber.SessionStore.Get(c)
	if err != nil {
		return err
	}
	if auth := session.Get("authenticated"); auth == nil {
		session.Set("authenticated", false)
	}
	session.Save()

	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("JWT_SECRET")), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	name := "Guest"
	user := getUserByEmail(claims.Issuer)
	if user != nil {
		name = user.DisplayName
	}
	return c.Render("index", fiber.Map{
		"Title":  "Hi, Planet!",
		"Forums": forums,
		"User":   name,
	})
}

// Get About
func GetAbout(c *fiber.Ctx) error {
	session, err := goth_fiber.SessionStore.Get(c)
	if err != nil {
		return err
	}
	fmt.Printf("======= GetAbout session.ID:%v\n", session.ID())
	fmt.Printf("======= GetAbout session.auth:%v\n", session.Get("authenticated"))
	test := session.Get("user")
	fmt.Printf("======= GetAbout user:%T\n", test)
	db := database.DBConn
	var forums []models.Forum
	db.Find(&forums)
	// return c.JSON(forums)
	return c.Render("about", fiber.Map{
		"Title": "About",
	})
}

// Get CreateForum
func GetCreateForum(c *fiber.Ctx) error {
	session, err := goth_fiber.SessionStore.Get(c)
	if err != nil {
		log.Fatal(err)
		return c.Redirect("/login")
	}

	user := session.Get("user")
	fmt.Println("user in GetCreateForum:", user)
	// return c.JSON(forums)
	return c.Render("create_forum", fiber.Map{
		"Title": "Create Forum",
	})
}

// Post CreateForum
func PostCreateForum(c *fiber.Ctx) error {
	session, err := goth_fiber.SessionStore.Get(c)
	if err != nil {
		log.Fatal(err)
		return c.Redirect("/login")
	}

	user := session.Get("user")
	fmt.Println("user in PostCreateForum:", user)
	db := database.DBConn
	forum := new(models.Forum)

	if err := c.BodyParser(forum); err != nil {
		return c.Status(503).SendString("The Error: " + err.Error())
	}
	// forum.User = c.
	db.Create(&forum)
	return c.Redirect("/")
	// return c.Render("index", fiber.Map{
	// 	"Title": "TFG",
	// })
}
