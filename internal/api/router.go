package api

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/admin"
	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/api/image"
	"github.com/JSYoo5B/SandStack/internal/api/placement"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	"github.com/go-chi/chi/v5"
)

func NewRouter(cfg config.Config) http.Handler {
	router := chi.NewRouter()
	router.Use(requestID)
	identityHandler := identity.NewHandler(cfg)

	router.Mount("/_sandstack", admin.NewRouter())

	router.Get("/identity", identityHandler.Discovery())
	router.Get("/identity/", identityHandler.Discovery())
	router.Mount("/identity/v3", identityHandler.Router())

	router.Mount("/compute/v2.1", compute.NewRouter(cfg))
	router.Mount("/image/v2", image.NewRouter(cfg))
	router.Mount("/placement", placement.NewRouter(cfg))

	return router
}

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Openstack-Request-Id", "req-"+idgen.RandomHex(16))
		next.ServeHTTP(w, r)
	})
}
