package main

import (
	"context"
	"log"

	"github.com/marcelofabianov/my-cash/config"
	"github.com/marcelofabianov/my-cash/pkg/database"
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

	ctx := context.Background()

	// Database
	db, err := database.Connect(ctx, cfg.Db)
	if err != nil {
		logger.Fatal("error connecting to database", logger.FieldError(err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal("error closing database connection")
		}
	}()

	logger.Info("starting application")
}
