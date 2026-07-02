package volume

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/go-chi/chi/v5"
)

func (h Handler) actionVolume(w http.ResponseWriter, r *http.Request) {
	var body map[string]json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	volumeID := chi.URLParam(r, "volume_id")
	for action, payload := range body {
		if h.handleVolumeAction(w, volumeID, action, payload) {
			return
		}
	}

	respond.Error(w, http.StatusBadRequest, "unsupported volume action")
}

func (h Handler) handleVolumeAction(
	w http.ResponseWriter,
	volumeID string,
	action string,
	payload json.RawMessage,
) bool {
	switch action {
	case "os-attach":
		var request attachVolumeRequest
		if !decodeVolumeAction(w, payload, &request) {
			return true
		}
		h.respondVolumeAction(w, h.service.Attach(volumeID, appvolume.AttachVolume{
			InstanceUUID: request.InstanceUUID,
			HostName:     request.HostName,
			MountPoint:   request.MountPoint,
			Mode:         request.Mode,
		}))
		return true
	case "os-begin_detaching":
		h.respondVolumeAction(w, h.service.BeginDetaching(volumeID))
		return true
	case "os-detach":
		var request detachVolumeRequest
		if !decodeVolumeAction(w, payload, &request) {
			return true
		}
		h.respondVolumeAction(w, h.service.Detach(volumeID, appvolume.DetachVolume{
			AttachmentID: request.AttachmentID,
		}))
		return true
	case "os-reserve":
		h.respondVolumeAction(w, h.service.Reserve(volumeID))
		return true
	case "os-unreserve":
		h.respondVolumeAction(w, h.service.Unreserve(volumeID))
		return true
	case "os-extend":
		var request extendVolumeRequest
		if !decodeVolumeAction(w, payload, &request) {
			return true
		}
		h.respondVolumeAction(w, h.service.ExtendSize(volumeID, request.NewSize))
		return true
	case "os-reset_status":
		var request resetVolumeStatusRequest
		if !decodeVolumeAction(w, payload, &request) {
			return true
		}
		h.respondVolumeAction(w, h.service.ResetStatus(
			volumeID,
			appvolume.ResetVolumeStatus{
				Status:       request.Status,
				AttachStatus: request.AttachStatus,
			},
		))
		return true
	default:
		return false
	}
}

func (h Handler) respondVolumeAction(w http.ResponseWriter, err error) {
	if errors.Is(err, appvolume.ErrVolumeNotFound) {
		respond.Error(w, http.StatusNotFound, "volume not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "volume action failed")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func decodeVolumeAction(
	w http.ResponseWriter,
	payload json.RawMessage,
	target any,
) bool {
	if err := json.Unmarshal(payload, target); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid volume action body")
		return false
	}

	return true
}
