package error

import (
	"log"

	"github.com/marcelofabianov/my-cash/config"
	"github.com/marcelofabianov/my-cash/pkg/logger"
)

func InitLogger() *logger.Logger {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	logger, err := logger.NewLogger(cfg.Log)
	if err != nil {
		log.Fatalf("error creating logger: %v", err)
	}

	return logger
}
