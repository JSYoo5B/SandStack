package placement

import (
	"net/http"
	"strings"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) versions(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, versionsResponse{
		Versions: []versionDocument{
			{
				ID:     "v1.0",
				Status: "CURRENT",
				Links: []versionLink{
					{
						Rel:  "self",
						Href: h.baseURL(r) + "/placement",
					},
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
