package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	param := chi.URLParam(r, "id")

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		return 0, err
	}

	return int64(id), nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<30) // 1MB

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	return dec.Decode(data)
}

func (app *application) writeJSONError(w http.ResponseWriter, status int, message string) {
	err := app.writeJSON(w, status, map[string]string{"error": message})
	log.Println(err)
}
