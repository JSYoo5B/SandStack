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

func (h Handler) getServerMetadatum(w http.ResponseWriter, r *http.Request) {
	meta, err := h.service.ServerMetadatum(
		chi.URLParam(r, "server_id"),
		chi.URLParam(r, "key"),
	)
	if errors.Is(err, appcompute.ErrServerNotFound) ||
		errors.Is(err, appcompute.ErrMetadatumNotFound) {
		respond.Error(w, http.StatusNotFound, "metadata not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "metadata lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, metadatumResponse{
		Meta: meta,
	})
}

func (h Handler) setServerMetadatum(w http.ResponseWriter, r *http.Request) {
	var request metadatumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	key := chi.URLParam(r, "key")
	value, ok := request.Meta[key]
	if !ok {
		respond.Error(w, http.StatusBadRequest, "metadata key mismatch")
		return
	}

	meta, err := h.service.SetServerMetadatum(
		chi.URLParam(r, "server_id"),
		key,
		value,
	)
	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "metadata update failed")
		return
	}

	respond.JSON(w, http.StatusOK, metadatumResponse{
		Meta: meta,
	})
}

func (h Handler) deleteServerMetadatum(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteServerMetadatum(
		chi.URLParam(r, "server_id"),
		chi.URLParam(r, "key"),
	)
	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "metadata delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
