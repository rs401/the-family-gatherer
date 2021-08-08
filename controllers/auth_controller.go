package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
	"github.com/shareed2k/goth_fiber"
)

func getUserByEmail(e string) *models.User {
	db := database.DBConn
	var user models.User
	db.Where("email = ?", e).Find(&user)
	if user.Email == e {
		return &user
	}
	return nil
}

func AuthCallback(c *fiber.Ctx) error {
	returnedUser, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		log.Fatal(err)
	}
	db := database.DBConn

	theUser := getUserByEmail(returnedUser.Email)
	token := returnedUser.AccessToken
	c.Set("token", token)
	c.Locals("token", token)

	if theUser == nil {
		newUser := new(models.User)
		newUser.Email = returnedUser.Email
		if returnedUser.NickName != "" {
			newUser.DisplayName = returnedUser.NickName
		} else if returnedUser.Name != "" {
			newUser.DisplayName = returnedUser.Name
		} else {
			newUser.DisplayName = returnedUser.Email
		}
		db.Create(newUser)
		theUser = newUser
	}
	session, err := goth_fiber.SessionStore.Get(c)
	if err != nil {
		return err
	}

	fmt.Printf("======== Session.ID: %v\n", session.ID())
	fmt.Printf("======== returnedUser.AccessToken: %v\n", returnedUser.AccessToken)
	session.Set("user", theUser)
	session.Set("token", token)
	session.Save()

	return c.Redirect("/")
}

func Logout(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		log.Fatal(err)
	}

	return c.Redirect("/")
}
