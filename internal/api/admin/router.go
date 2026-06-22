package admin

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/app/requestlog"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	return NewRouterWithReset(func() {})
}

func NewRouterWithReset(reset func()) http.Handler {
	return NewRouterWithState(reset, requestlog.NewService())
}

func NewRouterWithState(
	reset func(),
	requests *requestlog.Service,
) http.Handler {
	handler := Handler{
		reset:    reset,
		requests: requests,
	}
	router := chi.NewRouter()
	router.Get("/health", status)
	router.Get("/ready", status)
	router.Post("/reset", handler.resetState)
	router.Get("/requests", handler.listRequests)

	return router
}

type Handler struct {
	reset    func()
	requests *requestlog.Service
}
