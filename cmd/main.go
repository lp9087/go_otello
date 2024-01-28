package main

import (
	_ "github.com/lib/pq"
	"github.com/lp9087/go_otello_dashboard_api/config"
	"github.com/lp9087/go_otello_dashboard_api/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
