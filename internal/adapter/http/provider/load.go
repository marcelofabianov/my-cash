package provider

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/my-cash/pkg/logger"
)

func Load(app *fiber.App, db *sql.DB, logger *logger.Logger) *fiber.App {
	//app.Use(ServiceProvider(db, *logger))

	return app
}
