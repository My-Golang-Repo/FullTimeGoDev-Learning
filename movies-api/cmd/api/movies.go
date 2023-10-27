package main

import (
	"encoding/json"
	"fmt"
	"github.com/PorcoGalliard/GreenLight-Movie-API/internal/data"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Year:      2021,
		Runtime:   102,
		Genres:    []string{"romance", "action"},
		Version:   1,
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
