package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type PemainBola struct {
	Id string `json:"id"`
	Nama string `json:"nama"`
	Umur int `json:"umur"`
	Gender bool `json:"gender"`
}


func UserRoute(app fiber.Router) {
	app.Get("", func(c *fiber.Ctx) error {
		return c.SendString("Halo dari latihan fiber rafi")
	})

	// penggunaan body
	app.Post("", func (c *fiber.Ctx) error {
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
	app.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.JSON(
			map[string]any{
				"messages": fmt.Sprintf("Selamat datang %s", name),
			})
	})

	// penggunaan query params
	app.Get("", func(c *fiber.Ctx) error {
		umur := c.Query("umur")
		return c.JSON(
			map[string]any{
				"umur": umur,
			},
		)
	})

}