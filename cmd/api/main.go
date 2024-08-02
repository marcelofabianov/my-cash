package main

import (
	"log"

	"github.com/marcelofabianov/my-cash/config"
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

	logger.Info("starting application")
}
