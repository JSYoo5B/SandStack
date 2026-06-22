package volume

import (
	"net/http"
	"strings"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/go-chi/chi/v5"
)

func (h Handler) version(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")

	respond.JSON(w, http.StatusOK, versionResponse{
		Version: versionDocument{
			ID:      "v3.0",
			Status:  "CURRENT",
			Version: "3.0",
			Links: []versionLink{
				{
					Rel:  "self",
					Href: h.baseURL(r) + "/volume/v3/" + projectID,
				},
			},
		},
	})
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
