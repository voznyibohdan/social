package main

import (
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	if err != nil {
		app.writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
