package identity

import (
	"net/http"
	"strings"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	storeidentity "github.com/JSYoo5B/SandStack/internal/store/identity"
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
	return NewHandlerWithService(
		cfg,
		appidentity.NewServiceWithRepositories(
			cfg,
			appidentity.Repositories{
				Users:     storeidentity.NewMemoryUserRepository(),
				Projects:  storeidentity.NewMemoryProjectRepository(),
				Roles:     storeidentity.NewMemoryRoleRepository(),
				Tokens:    storeidentity.NewMemoryTokenRepository(),
				Services:  storeidentity.NewMemoryServiceRepository(),
				Endpoints: storeidentity.NewMemoryEndpointRepository(),
			},
		),
	)
}

func NewHandlerWithService(
	cfg config.Config,
	service appidentity.Service,
) Handler {
	return Handler{
		service: service,
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
	router.Get("/auth/tokens", h.getToken)
	router.Head("/auth/tokens", h.validateToken)
	router.Delete("/auth/tokens", h.revokeToken)
	router.Get("/projects", h.listProjects)
	router.Get("/projects/{project_id}", h.getProject)
	router.Get("/users", h.listUsers)
	router.Get("/users/{user_id}", h.getUser)
	router.Get("/roles", h.listRoles)
	router.Get("/roles/{role_id}", h.getRole)
	router.Get("/services", h.listServices)
	router.Get("/services/{service_id}", h.getService)
	router.Get("/endpoints", h.listEndpoints)
	router.Get("/endpoints/{endpoint_id}", h.getEndpoint)

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
