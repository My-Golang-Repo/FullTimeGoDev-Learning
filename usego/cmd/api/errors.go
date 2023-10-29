package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, message string, status int) {
	env := envelope{"error": message}

	if err := app.writeJSON(w, env, status, nil); err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}
