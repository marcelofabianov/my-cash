package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//...

		return c.Next()
	}
}
