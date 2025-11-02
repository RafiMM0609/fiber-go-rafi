package main

import (
	. "fiber-go/user"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)
func main() {
	app := fiber.New()

	// penggunaan middleware
	app.Use(LoggerMiddleware)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, OPTION, PUT, DELETE",
	}))

	userApp := app.Group("/v1/user")
	UserRoute(userApp)

	log.Fatal(app.Listen(":8000"))
}