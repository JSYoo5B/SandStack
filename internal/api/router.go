package api

import (
	"net/http"
	"strings"

	"github.com/JSYoo5B/SandStack/internal/api/admin"
	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/api/image"
	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/api/placement"
	"github.com/JSYoo5B/SandStack/internal/api/volume"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/JSYoo5B/SandStack/internal/app/requestlog"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	"github.com/go-chi/chi/v5"
)

func NewRouter(cfg config.Config) http.Handler {
	router := chi.NewRouter()
	router.Use(requestID)
	requests := requestlog.NewService()
	router.Use(recordRequests(requests))
	identityHandler := identity.NewHandler(cfg)
	computeService := appcompute.NewService()
	imageService := appimage.NewService()
	networkService := appnetwork.NewService()
	volumeService := appvolume.NewService()

	router.Mount("/_sandstack", admin.NewRouterWithState(func() {
		computeService.Reset()
		imageService.Reset()
		networkService.Reset()
		volumeService.Reset()
		requests.Reset()
	}, requests))

	router.Get("/identity", identityHandler.Discovery())
	router.Get("/identity/", identityHandler.Discovery())
	router.Mount("/identity/v3", identityHandler.Router())

	router.Mount("/compute/v2.1", compute.NewRouterWithService(cfg, computeService))
	router.Mount("/image/v2", image.NewRouterWithService(cfg, imageService))
	router.Mount("/network/v2.0", network.NewRouterWithService(cfg, networkService))
	router.Mount("/placement", placement.NewRouter(cfg))
	router.Mount("/volume/v3", volume.NewRouterWithService(cfg, volumeService))

	return router
}

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Openstack-Request-Id", "req-"+idgen.RandomHex(16))
		next.ServeHTTP(w, r)
	})
}

func recordRequests(requests *requestlog.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			recorder := &statusRecorder{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			next.ServeHTTP(recorder, r)

			if strings.HasPrefix(r.URL.Path, "/_sandstack") {
				return
			}

			requests.Add(requestlog.Record{
				ID:     w.Header().Get("X-Openstack-Request-Id"),
				Method: r.Method,
				Path:   r.URL.Path,
				Status: recorder.status,
			})
		})
	}
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}
