package admin

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/health", status)
	router.Get("/ready", status)

	return router
}
