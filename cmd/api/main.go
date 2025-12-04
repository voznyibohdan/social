package main

import (
	"log"

	"github.com/voznyibohdan/social/internal/storage"
)

func main() {
	app := &application{
		config:  loadConfig(),
		storage: storage.NewPostgresStorage(nil),
	}

	log.Fatal(app.serve(app.mount()))
}
