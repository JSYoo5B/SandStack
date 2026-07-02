package volume

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listVolumes(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, volumeListResponse{
		Volumes: toVolumeDocuments(h.service.List()),
	})
}

func (h Handler) createVolume(w http.ResponseWriter, r *http.Request) {
	var request createVolumeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	volume := h.service.Create(request.createVolume())
	respond.JSON(w, http.StatusAccepted, volumeResponse{
		Volume: toVolumeDocument(volume),
	})
}

func (h Handler) getVolume(w http.ResponseWriter, r *http.Request) {
	volume, err := h.service.Get(chi.URLParam(r, "volume_id"))
	if errors.Is(err, appvolume.ErrVolumeNotFound) {
		respond.Error(w, http.StatusNotFound, "volume not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "volume lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, volumeResponse{
		Volume: toVolumeDocument(volume),
	})
}

func (h Handler) updateVolume(w http.ResponseWriter, r *http.Request) {
	var request updateVolumeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	volume, err := h.service.Update(
		chi.URLParam(r, "volume_id"),
		request.updateVolume(),
	)
	if errors.Is(err, appvolume.ErrVolumeNotFound) {
		respond.Error(w, http.StatusNotFound, "volume not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "volume update failed")
		return
	}

	respond.JSON(w, http.StatusOK, volumeResponse{
		Volume: toVolumeDocument(volume),
	})
}

func (h Handler) deleteVolume(w http.ResponseWriter, r *http.Request) {
	err := h.service.Delete(chi.URLParam(r, "volume_id"))
	if errors.Is(err, appvolume.ErrVolumeNotFound) {
		respond.Error(w, http.StatusNotFound, "volume not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "volume delete failed")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
