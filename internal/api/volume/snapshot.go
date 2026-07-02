package volume

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listSnapshots(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, snapshotListResponse{
		Snapshots: toSnapshotDocuments(h.service.ListSnapshots()),
	})
}

func (h Handler) createSnapshot(w http.ResponseWriter, r *http.Request) {
	var request createSnapshotRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	snapshot := h.service.CreateSnapshot(request.createSnapshot())
	respond.JSON(w, http.StatusAccepted, snapshotResponse{
		Snapshot: toSnapshotDocument(snapshot),
	})
}

func (h Handler) getSnapshot(w http.ResponseWriter, r *http.Request) {
	snapshot, err := h.service.GetSnapshot(chi.URLParam(r, "snapshot_id"))
	if errors.Is(err, appvolume.ErrSnapshotNotFound) {
		respond.Error(w, http.StatusNotFound, "snapshot not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "snapshot lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, snapshotResponse{
		Snapshot: toSnapshotDocument(snapshot),
	})
}

func (h Handler) deleteSnapshot(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteSnapshot(chi.URLParam(r, "snapshot_id"))
	if errors.Is(err, appvolume.ErrSnapshotNotFound) {
		respond.Error(w, http.StatusNotFound, "snapshot not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "snapshot delete failed")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
