package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/rs401/TFG/database"
	"github.com/rs401/TFG/routes"
)

func main() {
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

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "https://127.0.0.1:3000/auth/callback"),
	)

	database.InitDatabase()
	// setupRoutes(app)
	routes.SetupRoutes(app)

	// app.Listen(":3000")
	// ln, _ := net.Listen("tcp", (":" + os.Getenv("API_PORT")))

	// cer, _ := tls.LoadX509KeyPair("server.crt", "server.key")

	// ln = tls.NewListener(ln, &tls.Config{Certificates: []tls.Certificate{cer}})

	// app.Listener(ln)
	app.ListenTLS(":"+os.Getenv("API_PORT"), "./server.crt", "./server.key")
}
