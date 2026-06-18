package identity

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) discovery(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, versionsResponse{
		Versions: versionValues{
			Values: []versionDocument{
				toVersionDocument(h.service.Version(h.baseURL(r))),
			},
		},
	})
}

func (h Handler) version(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, versionResponse{
		Version: toVersionDocument(h.service.Version(h.baseURL(r))),
	})
}
