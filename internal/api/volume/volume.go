package volume

import (
	"encoding/json"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
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
