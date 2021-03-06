package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs401/TFG/tfg-backend/database"
	"github.com/rs401/TFG/tfg-backend/routes"
)

func main() {
	database.InitDatabase()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen("127.0.0.1:8000")
}
