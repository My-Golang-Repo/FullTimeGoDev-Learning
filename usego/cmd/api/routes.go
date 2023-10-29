package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", app.homeHandler)
	router.Get("/v1/contact/{id}", app.contactHandler)
	router.Get("/v1/faq", app.faqHandler)

	return router
}
