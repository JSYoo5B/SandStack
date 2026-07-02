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

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewHandler(cfg config.Config) Handler {
	return Handler{
		service: appidentity.NewService(cfg),
		config:  cfg,
	}
}

func (h Handler) Discovery() http.HandlerFunc {
	return h.discovery
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/", h.version)
	router.Post("/auth/tokens", h.createToken)
	router.Get("/projects", h.listProjects)
	router.Get("/projects/{project_id}", h.getProject)
	router.Get("/users", h.listUsers)
	router.Get("/users/{user_id}", h.getUser)

	return router
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
