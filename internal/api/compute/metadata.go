package compute

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
)

func (h Handler) getServerMetadata(w http.ResponseWriter, r *http.Request) {
	metadata, err := h.service.ServerMetadata(chi.URLParam(r, "server_id"))
	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "metadata lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, metadataResponse{
		Metadata: metadata,
	})
}

func (h Handler) resetServerMetadata(w http.ResponseWriter, r *http.Request) {
	var request metadataRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	metadata, err := h.service.ResetServerMetadata(
		chi.URLParam(r, "server_id"),
		request.Metadata,
	)
	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "metadata reset failed")
		return
	}

	respond.JSON(w, http.StatusOK, metadataResponse{
		Metadata: metadata,
	})
}

func (h Handler) updateServerMetadata(w http.ResponseWriter, r *http.Request) {
	var request metadataRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	metadata, err := h.service.UpdateServerMetadata(
		chi.URLParam(r, "server_id"),
		request.Metadata,
	)
	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "metadata update failed")
		return
	}

	respond.JSON(w, http.StatusOK, metadataResponse{
		Metadata: metadata,
	})
}
