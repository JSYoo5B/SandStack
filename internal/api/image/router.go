package image

import (
	"net/http"

	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config  config.Config
	service *appimage.Service
}

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewHandler(cfg config.Config) Handler {
	return Handler{
		config:  cfg,
		service: appimage.NewService(),
	}
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/", h.versions)
	router.Get("/images", h.listImages)
	router.Post("/images", h.createImage)
	router.Get("/images/{image_id}", h.getImage)
	router.Delete("/images/{image_id}", h.deleteImage)

	return router
}
