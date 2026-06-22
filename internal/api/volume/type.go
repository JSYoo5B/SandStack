package volume

import (
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listVolumeTypes(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, volumeTypeListResponse{
		VolumeTypes: toVolumeTypeDocuments(h.service.ListVolumeTypes()),
	})
}

func (h Handler) getVolumeType(w http.ResponseWriter, r *http.Request) {
	volumeType, err := h.service.GetVolumeType(chi.URLParam(r, "type_id"))
	if errors.Is(err, appvolume.ErrVolumeTypeNotFound) {
		respond.Error(w, http.StatusNotFound, "volume type not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "volume type lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, volumeTypeResponse{
		VolumeType: toVolumeTypeDocument(volumeType),
	})
}
