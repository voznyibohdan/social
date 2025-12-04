package main

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
)

type config struct {
	Server serverConfig `envPrefix:"SERVER_"`
	DB     dbConfig     `envPrefix:"DB_"`
}

type serverConfig struct {
	Addr string `env:"PORT" default:":8080"`
}

type dbConfig struct {
	DSN          string        `env:"DSN" default:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
	MaxOpenConns int           `env:"MAX_OPEN_CONNS" default:"25"`
	MaxIdleConns int           `env:"MAX_IDLE_CONNS" default:"25"`
	MaxIdleTime  time.Duration `env:"MAX_IDLE_TIME" default:"15m"`
}

func loadConfig() *config {
	var cfg config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal("Error parsing env variables to config")
	}

	return &cfg
}
