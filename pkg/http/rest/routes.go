// Package rest implements a simple REST API for the
// application including routes and handlers.
package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/foo/{id}", app.fetchFoo)
	router.Post("/foo", app.addFoo)

	return router
}
