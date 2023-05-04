package main

import (
	"log"

	"github.com/joopinho/rp-scarper/configs"
	"github.com/joopinho/rp-scarper/internal/app"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	cfg := &configs.ServerConfig{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err)
	}
	app := app.NewEnrichApplication(cfg)
	app.Serve()
}
