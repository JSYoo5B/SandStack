package image

import (
	"encoding/json"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listImages(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, imageListResponse{
		Images: toImageDocuments(h.service.List()),
	})
}

func (h Handler) createImage(w http.ResponseWriter, r *http.Request) {
	var request createImageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	image := h.service.Create(request.createImage())
	respond.JSON(w, http.StatusCreated, toImageDocument(image))
}
