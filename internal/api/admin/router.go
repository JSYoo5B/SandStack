package admin

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/app/fault"
	"github.com/JSYoo5B/SandStack/internal/app/requestlog"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	return NewRouterWithReset(func() {})
}

func NewRouterWithReset(reset func()) http.Handler {
	return NewRouterWithState(reset, requestlog.NewService(), fault.NewService())
}

func NewRouterWithState(
	reset func(),
	requests *requestlog.Service,
	faults *fault.Service,
) http.Handler {
	handler := Handler{
		reset:    reset,
		requests: requests,
		faults:   faults,
	}
	router := chi.NewRouter()
	router.Get("/health", status)
	router.Get("/ready", status)
	router.Post("/reset", handler.resetState)
	router.Get("/requests", handler.listRequests)
	router.Get("/faults", handler.listFaults)
	router.Post("/faults", handler.createFault)
	router.Post("/faults/{fault_id}/enable", handler.enableFault)
	router.Post("/faults/{fault_id}/disable", handler.disableFault)

	return router
}

type Handler struct {
	reset    func()
	requests *requestlog.Service
	faults   *fault.Service
}
