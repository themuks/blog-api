package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/themuks/blog-api/routes"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	routes.PublicRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("API_PORT")))
}
