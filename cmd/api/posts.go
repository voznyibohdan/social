package main

import (
	"log"
	"net/http"

	"github.com/voznyibohdan/social/internal/storage"
)

type createPostDto struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var userID int64 = 1

	var dto createPostDto
	if err := app.readJSON(w, r, &dto); err != nil {
		app.writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &storage.Post{
		Title:   dto.Title,
		Content: dto.Content,
		Tags:    dto.Tags,
		UserID:  userID,
	}

	log.Printf("post: %+v\n", post)

	if err := app.storage.Posts.Create(r.Context(), post); err != nil {
		app.writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := app.writeJSON(w, http.StatusCreated, post); err != nil {
		app.writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
