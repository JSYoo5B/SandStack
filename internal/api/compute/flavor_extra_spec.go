package compute

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listFlavorExtraSpecs(w http.ResponseWriter, r *http.Request) {
	extraSpecs, err := h.service.ListFlavorExtraSpecs(
		chi.URLParam(r, "flavor_id"),
	)
	if errors.Is(err, appcompute.ErrFlavorNotFound) {
		respond.Error(w, http.StatusNotFound, "flavor not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "flavor lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, flavorExtraSpecListResponse{
		ExtraSpecs: extraSpecs,
	})
}

func (h Handler) getFlavorExtraSpec(w http.ResponseWriter, r *http.Request) {
	extraSpec, err := h.service.GetFlavorExtraSpec(
		chi.URLParam(r, "flavor_id"),
		chi.URLParam(r, "key"),
	)
	if errors.Is(err, appcompute.ErrFlavorNotFound) {
		respond.Error(w, http.StatusNotFound, "flavor extra spec not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "flavor lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, extraSpec)
}

func (h Handler) createFlavorExtraSpecs(w http.ResponseWriter, r *http.Request) {
	var request createFlavorExtraSpecsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid flavor extra spec request")
		return
	}

	extraSpecs, err := h.service.CreateFlavorExtraSpecs(
		chi.URLParam(r, "flavor_id"),
		request.ExtraSpecs,
	)
	if errors.Is(err, appcompute.ErrFlavorNotFound) {
		respond.Error(w, http.StatusNotFound, "flavor not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "flavor update failed")
		return
	}

	respond.JSON(w, http.StatusOK, flavorExtraSpecListResponse{
		ExtraSpecs: extraSpecs,
	})
}

func (h Handler) updateFlavorExtraSpec(w http.ResponseWriter, r *http.Request) {
	var request map[string]string
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid flavor extra spec request")
		return
	}

	extraSpec, err := h.service.UpdateFlavorExtraSpec(
		chi.URLParam(r, "flavor_id"),
		chi.URLParam(r, "key"),
		request,
	)
	if errors.Is(err, appcompute.ErrFlavorNotFound) {
		respond.Error(w, http.StatusNotFound, "flavor extra spec not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "flavor update failed")
		return
	}

	respond.JSON(w, http.StatusOK, extraSpec)
}

func (h Handler) deleteFlavorExtraSpec(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteFlavorExtraSpec(
		chi.URLParam(r, "flavor_id"),
		chi.URLParam(r, "key"),
	)
	if errors.Is(err, appcompute.ErrFlavorNotFound) {
		respond.Error(w, http.StatusNotFound, "flavor not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "flavor update failed")
		return
	}

	respond.JSON(w, http.StatusOK, map[string]any{})
}
