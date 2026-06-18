package identity

import (
	"net/http"
	"strings"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service appidentity.Service
	config  config.Config
}

func Discovery(cfg config.Config) http.HandlerFunc {
	return NewHandler(cfg).discovery
}

func NewRouter(cfg config.Config) http.Handler {
	handler := NewHandler(cfg)

	router := chi.NewRouter()
	router.Get("/", handler.version)
	router.Post("/auth/tokens", handler.createToken)

	return router
}

func NewHandler(cfg config.Config) Handler {
	return Handler{
		service: appidentity.NewService(cfg),
		config:  cfg,
	}
}

func (h Handler) baseURL(r *http.Request) string {
	if h.config.PublicURL != "" {
		return strings.TrimRight(h.config.PublicURL, "/")
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}
