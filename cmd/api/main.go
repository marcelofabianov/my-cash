package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/my-cash/config"
	"github.com/marcelofabianov/my-cash/internal/adapter/http"
	"github.com/marcelofabianov/my-cash/internal/adapter/http/middleware"
	"github.com/marcelofabianov/my-cash/pkg/logger"
)

func main() {
	// Config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Logger
	logger, err := logger.NewLogger(cfg.Log)
	if err != nil {
		log.Fatalf("error creating logger: %v", err)
	}
	defer logger.Close()

	// API Server
	app := fiber.New()
	app = middleware.Load(app, logger, &cfg.Api)
	app = http.Route(app)

	addr := fmt.Sprintf("%s:%s", cfg.Api.Host, cfg.Api.Port)

	if err := app.Listen(addr); err != nil {
		logger.Fatal("error starting API server", logger.FieldError(err))
	}

	logger.Info("API server started", logger.Field("address", addr))
}
