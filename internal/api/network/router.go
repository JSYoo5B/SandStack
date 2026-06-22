package network

import (
	"net/http"

	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config  config.Config
	service *appnetwork.Service
}

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewHandler(cfg config.Config) Handler {
	return Handler{
		config:  cfg,
		service: appnetwork.NewService(),
	}
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/", h.versions)
	router.Get("/networks", h.listNetworks)
	router.Post("/networks", h.createNetwork)
	router.Get("/networks/{network_id}", h.getNetwork)
	router.Delete("/networks/{network_id}", h.deleteNetwork)
	router.Get("/subnets", h.listSubnets)

	return router
}
