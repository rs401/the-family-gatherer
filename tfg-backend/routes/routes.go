package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/tfg-backend/controllers"
)

func Setup(app *fiber.App) {
	// api
	api := app.Group("/api")
	// Auth
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Get("/user", controllers.User)
	api.Post("/logout", controllers.Logout)

	// Forum routes
	api.Get("/forum", controllers.GetForums)
	api.Get("/forum/:id", controllers.GetForum)
	api.Post("/forum", controllers.NewForum)
	api.Put("/forum/:id", controllers.UpdateForum)
	api.Delete("/forum/:id", controllers.DeleteForum)

	// Thread routes
	api.Get("/:fid/thread", controllers.GetThreads)
	api.Get("/thread/:id", controllers.GetThread)
	api.Post("/:fid/thread", controllers.NewThread)
	api.Put("/thread/:id", controllers.UpdateThread)
	api.Delete("/thread/:id", controllers.DeleteThread)

	// Post routes
	api.Get("/:tid/post", controllers.GetPosts)
	api.Get("/post/:id", controllers.GetPost)
	api.Post("/:tid/post", controllers.NewPost)
	api.Put("/post/:id", controllers.UpdatePost)
	api.Delete("/post/:id", controllers.DeletePost)

}
