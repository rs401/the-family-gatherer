package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

func AuthCallback(c *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendString(user.Email)
}

func Logout(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		log.Fatal(err)
	}

	return c.SendString("logout")
}
