package controllers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rs401/TFG/config"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/models"
	"github.com/shareed2k/goth_fiber"
	"golang.org/x/crypto/bcrypt"
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

func Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

func DoLogin(c *fiber.Ctx) error {
	email := c.FormValue("email")
	email = strings.TrimSpace(email)
	if email == "" {
		fmt.Println("Login Email was empty string")
		return nil
	}
	password := c.FormValue("password")

	user := getUserByEmail(email)
	if user == nil {
		// User doesn't exist
		fmt.Println("User doesn't exist.")
		return c.Redirect("/register")
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		// password and hash didn't match
		fmt.Printf("Hash and pass != : %s\n", err.Error())
		// Flash wrong password
		return c.Redirect("/login")
	}

	// Set cookie
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(config.Config("JWT_SECRET")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Redirect("/")
}

func Register(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	})
}

func DoRegister(c *fiber.Ctx) error {
	name := c.FormValue("displayname")
	email := c.FormValue("email")
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	pw1 := c.FormValue("password1")
	pw2 := c.FormValue("password2")

	// If empty name or email
	if name == "" || email == "" {
		fmt.Println("Displayname or Email were empty string.")
		// figure out how to flash validation
		return nil
	}
	// If passwords do not match
	if pw1 != pw2 {
		fmt.Println("Passwords did not match")
		// figure out how to flash validation
		return nil
	}
	// passwords match bcrypt it
	password, err := bcrypt.GenerateFromPassword([]byte(pw1), bcrypt.DefaultCost)

	if err != nil {
		fmt.Printf("Error hashing passsword: %s\n", err.Error())
	}

	db := database.DBConn
	user := models.User{
		DisplayName: name,
		Email:       email,
		Password:    password,
	}
	db.Create(&user)

	return c.Redirect("/login")
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
