package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/voznyibohdan/social/internal/storage"
)

type createPostDto struct {
	Title   string   `json:"title" validate:"required,max=50,min=3"`
	Content string   `json:"content" validate:"required,max=150,min=20"`
	Tags    []string `json:"tags" validate:"required,max=20,min=1"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var userID int64 = 1

	var dto createPostDto
	if err := app.readJSON(w, r, &dto); err != nil {
		app.writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := app.validate.Struct(dto)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fmt.Println(e.Translate(app.trans))
		}
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

func (app *application) getPostByID(w http.ResponseWriter, r *http.Request) {
	postID, err := app.readIDParam(r)
	if err != nil {
		app.writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := app.storage.Posts.GetOneByID(r.Context(), postID)
	if err != nil {
		app.writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := app.writeJSON(w, http.StatusCreated, post); err != nil {
		app.writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
