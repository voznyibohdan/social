package main

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type config struct {
	Server serverConfig `envPrefix:"SERVER_"`
}

type serverConfig struct {
	Addr string `env:"PORT" default:":8080"`
}

func loadConfig() *config {
	var cfg config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal("Error parsing env variables to config")
	}

	return &cfg
}
