package image

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config config.Config
}

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewHandler(cfg config.Config) Handler {
	return Handler{config: cfg}
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/", h.versions)
	router.Get("/images", h.listImages)

	return router
}
