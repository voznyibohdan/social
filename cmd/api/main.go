package main

import (
	"log"

	"github.com/voznyibohdan/social/internal/db"
	"github.com/voznyibohdan/social/internal/storage"
)

func main() {
	cfg := loadConfig()

	database, err := db.OpenDB(cfg.DB.DSN, cfg.DB.MaxIdleTime, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns)
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()
	log.Println("database connection pool established")

	app := &application{
		config:  cfg,
		storage: storage.NewPostgresStorage(database),
		db:      database,
	}

	log.Fatal(app.serve(app.mount()))
}
