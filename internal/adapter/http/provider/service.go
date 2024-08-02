package provider

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/my-cash/pkg/logger"
)

func ServiceProvider(db *sql.DB, logger logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//....

		return c.Next()
	}
}
