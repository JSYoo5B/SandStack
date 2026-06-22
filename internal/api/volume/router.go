package volume

import (
	"net/http"

	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config  config.Config
	service *appvolume.Service
}

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewHandler(cfg config.Config) Handler {
	return Handler{
		config:  cfg,
		service: appvolume.NewService(),
	}
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/{project_id}", h.version)
	router.Get("/{project_id}/", h.version)
	router.Get("/{project_id}/volumes/detail", h.listVolumes)
	router.Post("/{project_id}/volumes", h.createVolume)

	return router
}
