package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type PemainBola struct {
	Id string `json:"id"`
	Nama string `json:"nama"`
	Umur int `json:"umur"`
	Gender bool `json:"gender"`
}

func main() {
	app := fiber.New()

	// penggunaan middleware
	app.Use(LoggerMiddleware)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, OPTION, PUT, DELETE",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Halo dari latihan fiber rafi")
	})

	// penggunaan body
	app.Post("/user", func (c *fiber.Ctx) error {
		pemain := new(PemainBola)

		if err := c.BodyParser(pemain); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Input json tidak valid",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Input Valid",
			"data": pemain,
		})
	})

	// penggunaan params
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.JSON(
			map[string]any{
				"messages": fmt.Sprintf("Selamat datang %s", name),
			})
	})

	// penggunaan query params
	app.Get("/user", func(c *fiber.Ctx) error {
		umur := c.Query("umur")
		return c.JSON(
			map[string]any{
				"umur": umur,
			},
		)
	})


	log.Fatal(app.Listen(":8000"))
}