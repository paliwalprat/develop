package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pratiksha/blogbackend/controller"
	"github.com/pratiksha/blogbackend/middleware"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticated)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/getAllPost", controller.AllPost)
	app.Get("/api/getPost/:id", controller.DetailPosts)
	app.Put("/api/updatePost/:id", controller.UpdatePosts)
	app.Get("/api/uniquePost", controller.UniquePost)
	app.Delete("/api/deletePost/:id", controller.DeletePost)
	app.Post("/api/uploadImage", controller.Upload)

}
