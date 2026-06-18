package api

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/admin"
	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	"github.com/go-chi/chi/v5"
)

func NewRouter(cfg config.Config) http.Handler {
	router := chi.NewRouter()
	router.Use(requestID)

	router.Mount("/_sandstack", admin.NewRouter())

	router.Get("/identity", identity.Discovery(cfg))
	router.Get("/identity/", identity.Discovery(cfg))
	router.Mount("/identity/v3", identity.NewRouter(cfg))

	return router
}

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Openstack-Request-Id", "req-"+idgen.RandomHex(16))
		next.ServeHTTP(w, r)
	})
}
