package main

import (
	"fmt"
	"log"

	"github.com/marcelofabianov/my-cash/config"
)

func main() {
	// Config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Println(cfg.Env)
}
