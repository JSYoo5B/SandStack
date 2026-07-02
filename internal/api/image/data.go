package image

import (
	"errors"
	"io"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
	"github.com/go-chi/chi/v5"
)

func (h Handler) uploadImageData(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid image data")
		return
	}

	err = h.service.UploadData(chi.URLParam(r, "image_id"), data)
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image data upload failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) stageImageData(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid image data")
		return
	}

	err = h.service.StageData(chi.URLParam(r, "image_id"), data)
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image data stage failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) importImageData(w http.ResponseWriter, r *http.Request) {
	err := h.service.ImportData(chi.URLParam(r, "image_id"))
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image import failed")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h Handler) getImportInfo(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, importInfoResponse{
		ImportMethods: importMethodsDocument{
			Description: "Import methods available.",
			Type:        "array",
			Value: []string{
				"glance-direct",
			},
		},
	})
}

func (h Handler) downloadImageData(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.DownloadData(chi.URLParam(r, "image_id"))
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image data not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image data download failed")
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(data)
}
