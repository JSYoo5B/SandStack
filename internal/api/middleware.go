package api

import (
	"net/http"
	"strings"

	"github.com/JSYoo5B/SandStack/internal/app/requestlog"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

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
