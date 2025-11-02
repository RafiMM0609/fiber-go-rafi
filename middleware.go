package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	startTime := time.Now()
	fmt.Printf("waktu: %s , method: %s , path: %s \n",
		startTime.Format("2006-01-02 15:04:05"),
		c.Method(),
		c.Path(),
	)
	err := c.Next()
	fmt.Printf(
		"request method %s pada %s selesai dieksekusi, selama %d ms\n",
		c.Method(), c.Path(), time.Since(startTime).Milliseconds(),
	)
	return err
}

// func AuthMiddleware(c *fiber.Ctx) error {
// 	token := c.Get("Authorization")
// 	if token != "" {
// 		return c.Next()
// 	}
// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"message": "Forbidden Access",
// 	})
// }