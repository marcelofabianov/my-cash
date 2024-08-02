package provider

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/my-cash/internal/container"
	"github.com/marcelofabianov/my-cash/internal/port/inbound"
	"github.com/marcelofabianov/my-cash/pkg/logger"
)

func ServiceProvider(db *sql.DB, logger logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := registerUserService(c, db, logger); err != nil {
			return err
		}

		return c.Next()
	}
}

func registerUserService(c *fiber.Ctx, db *sql.DB, logger logger.Logger) error {
	userC := container.NewUserContainer(db, logger)

	var userS inbound.UserService
	err := userC.Invoke(func(service inbound.UserService) {
		userS = service
	})
	if err != nil {
		return err
	}

	c.Locals("userService", userS)

	return nil
}
