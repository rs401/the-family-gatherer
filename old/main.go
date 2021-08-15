package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/rs401/TFG/config"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/routes"
	"github.com/shareed2k/goth_fiber"
)

func main() {

	app := Setup()

	app.ListenTLS(":"+config.Config("API_PORT"), "./server.crt", "./server.key")
}

func Setup() *fiber.App {
	// Load env vars
	// if err := godotenv.Load(".env"); err != nil {
	// 	panic("Error loading .env file")
	// }
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")
	app.Use(logger.New())
	app.Use(cors.New())

	sconfig := session.Config{
		Expiration:     24 * time.Hour,
		KeyLookup:      "header:session_id",
		CookieDomain:   "",
		CookiePath:     "",
		CookieSecure:   true,
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
	sessions := session.New(sconfig)
	// sessions := session.New()

	goth_fiber.SessionStore = sessions
	goth.UseProviders(
		google.New(config.Config("GOOGLE_CLIENT_ID"), config.Config("GOOGLE_CLIENT_SECRET"), "https://127.0.0.1:3000/authcallback"),
	)

	database.InitDatabase()
	routes.SetupRoutes(app)
	return app
}