package compute

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
			ID:      "v2.1",
			Status:  "CURRENT",
			Version: "2.1",
			Links: []versionLink{
				{
					Rel:  "self",
					Href: h.baseURL(r) + "/compute/v2.1/" + projectID,
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
