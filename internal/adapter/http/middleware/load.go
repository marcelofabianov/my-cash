package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/my-cash/config"
	"github.com/marcelofabianov/my-cash/pkg/logger"
)

func Load(app *fiber.App, logger *logger.Logger, cfg *config.ApiConfig) *fiber.App {
	app.Use(RecoverMiddleware())

	if cfg.LogRequests {
		app.Use(LoggingMiddleware(logger))
	}

	app.Use(CorsMiddleware())
	app.Use(RateLimitMiddleware())
	app.Use(AuthMiddleware())
	app.Use(AuthorizeMiddleware())

	return app
}
