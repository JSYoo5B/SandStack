package volume

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listExtraSpecs(w http.ResponseWriter, r *http.Request) {
	extraSpecs, err := h.service.ListExtraSpecs(chi.URLParam(r, "type_id"))
	if handleExtraSpecError(w, err, "extra specs lookup failed") {
		return
	}

	respond.JSON(w, http.StatusOK, extraSpecsResponse{
		ExtraSpecs: extraSpecs,
	})
}

func (h Handler) createExtraSpecs(w http.ResponseWriter, r *http.Request) {
	var request extraSpecsResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	extraSpecs, err := h.service.CreateExtraSpecs(
		chi.URLParam(r, "type_id"),
		request.ExtraSpecs,
	)
	if handleExtraSpecError(w, err, "extra specs create failed") {
		return
	}

	respond.JSON(w, http.StatusOK, extraSpecsResponse{
		ExtraSpecs: extraSpecs,
	})
}

func (h Handler) getExtraSpec(w http.ResponseWriter, r *http.Request) {
	extraSpec, err := h.service.GetExtraSpec(
		chi.URLParam(r, "type_id"),
		chi.URLParam(r, "key"),
	)
	if handleExtraSpecError(w, err, "extra spec lookup failed") {
		return
	}

	respond.JSON(w, http.StatusOK, extraSpec)
}

func (h Handler) updateExtraSpec(w http.ResponseWriter, r *http.Request) {
	var request map[string]string
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	extraSpec, err := h.service.UpdateExtraSpec(
		chi.URLParam(r, "type_id"),
		request,
	)
	if handleExtraSpecError(w, err, "extra spec update failed") {
		return
	}

	respond.JSON(w, http.StatusOK, extraSpec)
}

func (h Handler) deleteExtraSpec(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteExtraSpec(
		chi.URLParam(r, "type_id"),
		chi.URLParam(r, "key"),
	)
	if handleExtraSpecError(w, err, "extra spec delete failed") {
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func handleExtraSpecError(
	w http.ResponseWriter,
	err error,
	internalMessage string,
) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, appvolume.ErrVolumeTypeNotFound) {
		respond.Error(w, http.StatusNotFound, "volume type not found")
		return true
	}
	if errors.Is(err, appvolume.ErrExtraSpecNotFound) {
		respond.Error(w, http.StatusNotFound, "extra spec not found")
		return true
	}

	respond.Error(w, http.StatusInternalServerError, internalMessage)
	return true
}
