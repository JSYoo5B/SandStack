package volume

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listBackups(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, backupListResponse{
		Backups: toBackupDocuments(h.service.ListBackups()),
	})
}

func (h Handler) createBackup(w http.ResponseWriter, r *http.Request) {
	var request createBackupRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	backup := h.service.CreateBackup(request.createBackup())
	respond.JSON(w, http.StatusAccepted, backupResponse{
		Backup: toBackupDocument(backup),
	})
}

func (h Handler) getBackup(w http.ResponseWriter, r *http.Request) {
	backup, err := h.service.GetBackup(chi.URLParam(r, "backup_id"))
	if errors.Is(err, appvolume.ErrBackupNotFound) {
		respond.Error(w, http.StatusNotFound, "backup not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "backup lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, backupResponse{
		Backup: toBackupDocument(backup),
	})
}

func (h Handler) deleteBackup(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteBackup(chi.URLParam(r, "backup_id"))
	if errors.Is(err, appvolume.ErrBackupNotFound) {
		respond.Error(w, http.StatusNotFound, "backup not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "backup delete failed")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
