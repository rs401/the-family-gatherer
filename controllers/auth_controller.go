package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
	"github.com/shareed2k/goth_fiber"
)

func AuthCallback(c *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		log.Fatal(err)
	}
	db := database.DBConn
	newUser := new(models.User)
	newUser.Email = user.Email
	if user.NickName != "" {
		newUser.DisplayName = user.NickName
	} else if user.Name != "" {
		newUser.DisplayName = user.Name
	} else {
		newUser.DisplayName = user.Email
	}
	session, err := goth_fiber.SessionStore.Get(c)
	if err != nil {
		return err
	}
	session.Set("user", user)
	db.Create(newUser)

	return c.Redirect("/")
}

func Logout(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		log.Fatal(err)
	}

	return c.Redirect("/")
}
