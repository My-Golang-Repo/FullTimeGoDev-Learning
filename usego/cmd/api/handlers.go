package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.parseTemplate(w, "templates/home.gohtml", nil); err != nil {
		log.Printf("passing template %v", err)
		//http.Error(w, "something error with template", http.StatusInternalServerError)
		app.errorResponse(w, r, "There was something error while parsing template", http.StatusInternalServerError)
	}
}

func (app *application) contactHandler(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")
	_, err := strconv.Atoi(params)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.parseTemplate(w, "templates/contact.gohtml", nil); err != nil {
		log.Printf("passing template %v", err)
		app.errorResponse(w, r, "There was something error wjile parsing template", http.StatusInternalServerError)
	}
}

func (app *application) faqHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.parseTemplate(w, "templates/faq.gohtml", nil); err != nil {
		log.Printf("passing template %v", err)
		app.errorResponse(w, r, "There was something error with the template", http.StatusInternalServerError)
	}
}
