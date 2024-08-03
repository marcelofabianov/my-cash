package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/my-cash/internal/adapter/http/handler"
)

func Route(app *fiber.App) *fiber.App {
	api := app.Group("/api/v1")

	// Health
	api.Get("/health", handler.HealthCheckHandler)

	// Users
	api.Post("/users", handler.CreateUserHandler)

	return app
}
