package image

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
	"github.com/go-chi/chi/v5"
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

func (h Handler) getImage(w http.ResponseWriter, r *http.Request) {
	image, err := h.service.Get(chi.URLParam(r, "image_id"))
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, toImageDocument(image))
}

func (h Handler) deleteImage(w http.ResponseWriter, r *http.Request) {
	err := h.service.Delete(chi.URLParam(r, "image_id"))
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
