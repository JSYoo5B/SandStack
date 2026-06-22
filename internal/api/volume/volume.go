package volume

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listVolumes(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, volumeListResponse{
		Volumes: []volumeDocument{},
	})
}
