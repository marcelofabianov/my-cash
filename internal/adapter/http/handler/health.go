package handler

import "github.com/gofiber/fiber/v2"

func HealthCheckHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	return c.Status(200).JSON(fiber.Map{
		"status":  "OK",
		"message": "Service is healthy",
	})
}
