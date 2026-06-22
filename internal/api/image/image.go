package image

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listImages(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, imageListResponse{
		Images: []imageDocument{},
	})
}
