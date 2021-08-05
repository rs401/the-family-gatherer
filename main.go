package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/routes"
	"github.com/shareed2k/goth_fiber"
)

func main() {

	app := Setup()

	app.ListenTLS(":"+os.Getenv("API_PORT"), "./server.crt", "./server.key")
}

func Setup() *fiber.App {
	// Load env vars
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")
	app.Use(logger.New())
	// app.Use(cors.New())

	config := session.Config{
		Expiration:     24 * time.Hour,
		Storage:        nil,
		KeyLookup:      "cookie:session_id",
		CookieDomain:   "",
		CookiePath:     "",
		CookieSecure:   false,
		CookieHTTPOnly: false,
		// CookieSameSite: "",
	}
	// config := session.Config{
	//     Key:    "dinosaurus",       // default: "sessionid"
	//     Lookup: "header",           // default: "cookie"
	//     Domain: "google.com",       // default: ""
	//     Expires: 30 * time.Minutes, // default: 2 * time.Hour
	//     Secure:  true,              // default: false
	//  }

	// create session handler
	sessions := session.New(config)

	goth_fiber.SessionStore = sessions
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "https://127.0.0.1:3000/auth/callback"),
	)

	database.InitDatabase()
	routes.SetupRoutes(app)
	return app
}
