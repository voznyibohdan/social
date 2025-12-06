package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/voznyibohdan/social/internal/storage"
)

type application struct {
	config   *config
	storage  *storage.Storage
	db       *sql.DB
	validate *validator.Validate
	trans    ut.Translator
}

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheck)

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
			r.Get("/{id}", app.getPostByID)
		})
	})

	return router
}

func (app *application) serve(handler http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.Server.Addr,
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return srv.ListenAndServe()
}
