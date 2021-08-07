package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
	"github.com/shareed2k/goth_fiber"
	"gorm.io/gorm"
)

func getUserByEmail(e string) (*models.User, error) {
	db := database.DBConn
	var user models.User
	if err := db.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func AuthCallback(c *fiber.Ctx) error {
	returnedUser, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		log.Fatal(err)
	}
	db := database.DBConn

	theUser, err := getUserByEmail(returnedUser.Email)
	if err != nil {
		log.Fatal(err)
		return c.Redirect("/login")
	}
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
	session.Set("user", theUser)

	return c.Redirect("/")
}

func Logout(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		log.Fatal(err)
	}

	return c.Redirect("/")
}
