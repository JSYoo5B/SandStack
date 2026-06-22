package admin

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	return NewRouterWithReset(func() {})
}

func NewRouterWithReset(reset func()) http.Handler {
	handler := Handler{reset: reset}
	router := chi.NewRouter()
	router.Get("/health", status)
	router.Get("/ready", status)
	router.Post("/reset", handler.resetState)

	return router
}

type Handler struct {
	reset func()
}
