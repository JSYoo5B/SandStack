package compute

import (
	"net/http"

	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config  config.Config
	service *appcompute.Service
}

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewHandler(cfg config.Config) Handler {
	return Handler{
		config:  cfg,
		service: appcompute.NewService(),
	}
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/{project_id}", h.version)
	router.Get("/{project_id}/", h.version)
	router.Get("/{project_id}/flavors", h.listFlavors)
	router.Get("/{project_id}/flavors/detail", h.listFlavors)
	router.Get("/{project_id}/flavors/{flavor_id}", h.getFlavor)

	return router
}
