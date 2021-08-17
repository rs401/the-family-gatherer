package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rs401/TFG/tfg-backend/config"
	"github.com/rs401/TFG/tfg-backend/database"
	"github.com/rs401/TFG/tfg-backend/models"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = config.Config("JWT_SECRET")

func getUserByEmail(e string) *models.User {
	db := database.DBConn
	var user models.User
	db.Where("email = ?", e).Find(&user)
	if user.Email == e {
		return &user
	}
	return nil
}

func getUserById(i string) *models.User {
	db := database.DBConn
	var user models.User
	id, err := strconv.Atoi(i)
	if err != nil {
		fmt.Printf("getUserById failed to convert AtoI: %s\n", err.Error())
		return nil
	}
	db.Where("id = ?", id).Find(&user)
	if int(user.ID) == id {
		return &user
	}
	return nil
}

func getUserByJwt(cookie string) *models.User {

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil // unauthenticated
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user := getUserById(claims.Issuer)
	if user == nil {
		return nil // user not found
	}

	return user
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if err := getUserByEmail(data["email"]); err != nil {
		ret := make(map[string]string)
		ret["message"] = "Email already exists in our database."
		return c.JSON(ret)
	}
	// Check passwords match
	pass1 := data["password1"]
	pass2 := data["password2"]
	fmt.Println("pass1: ", pass1)
	fmt.Println("pass2: ", pass2)
	if pass1 == "" || pass1 != pass2 {
		ret := make(map[string]string)
		ret["message"] = "Passwords do not match."
		return c.JSON(ret)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(data["password1"]), 14)
	if err != nil {
		fmt.Println("=========== Houston, we have a problem")
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DBConn.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := getUserByEmail(data["email"])
	if user == nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "tfg",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Successfully logged in.",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("tfg")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user := getUserById(claims.Issuer)
	if user == nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "tfg",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
