package main

import (
	"github.com/akhmettolegen/translator/internal/app"
	"log"

	"github.com/akhmettolegen/translator/config"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
