package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/themuks/blog-api/app/controllers"
	"github.com/themuks/blog-api/middleware"
)

func PublicRoutes(router *fiber.App) {
	router.Use(middleware.Json)

	router.Static("/", "./public")

	api := router.Group("/api/v1")

	articles := api.Group("/articles")
	articles.Get("/", controllers.GetArticles)
	articles.Post("/", controllers.AddArticle)
	articles.Get("/:id", controllers.GetArticle)
	articles.Put("/:id", controllers.UpdateArticle)
	articles.Delete("/:id", controllers.DeleteArticle)
	articles.Get("/:id/comments", controllers.GetArticleComments)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    "404",
			"message": "404: Not Found",
		})
	})
}
