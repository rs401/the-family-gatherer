package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs401/TFG/controllers"
	"github.com/shareed2k/goth_fiber"
)

func SetupRoutes(app *fiber.App) {

	// User auth
	app.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	app.Get("/login", controllers.Login)
	app.Get("/register", controllers.Register)
	app.Get("/logout", controllers.Logout)
	app.Get("/authcallback", controllers.AuthCallback)

	app.Post("/login", controllers.DoLogin)
	app.Post("/register", controllers.DoRegister)

	// HTML Routes
	app.Get("/", controllers.GetIndex)
	app.Get("/about", controllers.GetAbout)
	app.Get("/create_forum", controllers.GetCreateForum)
	app.Post("/create_forum", controllers.PostCreateForum)

	// API Routes
	// Forum routes
	app.Get("/api/v1/forum", controllers.GetForums)
	app.Get("/api/v1/forum/:id", controllers.GetForum)
	app.Post("/api/v1/forum", controllers.NewForum)
	app.Put("/api/v1/forum/:id", controllers.UpdateForum)
	app.Delete("/api/v1/forum/:id", controllers.DeleteForum)

	// Thread routes
	app.Get("/api/v1/:fid/thread", controllers.GetThreads)
	app.Get("/api/v1/thread/:id", controllers.GetThread)
	app.Post("/api/v1/:fid/thread", controllers.NewThread)
	app.Put("/api/v1/thread/:id", controllers.UpdateThread)
	app.Delete("/api/v1/thread/:id", controllers.DeleteThread)

	// Post routes
	app.Get("/api/v1/post/:tid", controllers.GetPosts)
	app.Get("/api/v1/forum/:fid/thread/:tid/post/:id", controllers.GetPost)
	app.Post("/api/v1/forum/:fid/thread/:tid/post", controllers.NewPost)
	app.Put("/api/v1/forum/:fid/thread/:id", controllers.UpdatePost)
	app.Delete("/api/v1/forum/:fid/thread/:id", controllers.DeletePost)

}
