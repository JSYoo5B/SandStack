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

func NewRouterWithService(
	cfg config.Config,
	service *appvolume.Service,
) http.Handler {
	return NewHandlerWithService(cfg, service).Router()
}

func NewHandler(cfg config.Config) Handler {
	return NewHandlerWithService(cfg, appvolume.NewService())
}

func NewHandlerWithService(
	cfg config.Config,
	service *appvolume.Service,
) Handler {
	return Handler{
		config:  cfg,
		service: service,
	}
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/{project_id}", h.version)
	router.Get("/{project_id}/", h.version)
	router.Get("/{project_id}/volumes/detail", h.listVolumes)
	router.Post("/{project_id}/volumes", h.createVolume)
	router.Get("/{project_id}/volumes/{volume_id}", h.getVolume)
	router.Delete("/{project_id}/volumes/{volume_id}", h.deleteVolume)
	router.Get("/{project_id}/types", h.listVolumeTypes)
	router.Get("/{project_id}/types/{type_id}", h.getVolumeType)

	return router
}
