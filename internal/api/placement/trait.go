package placement

import (
	"encoding/json"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
	"github.com/go-chi/chi/v5"
)

func (h Handler) getTraits(w http.ResponseWriter, r *http.Request) {
	traits, err := h.service.GetTraits(
		chi.URLParam(r, "resource_provider_uuid"),
	)
	if handlePlacementError(w, err, "trait lookup failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toTraitsDocument(traits))
}

func (h Handler) updateTraits(w http.ResponseWriter, r *http.Request) {
	var request traitsDocument
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	traits, err := h.service.UpdateTraits(
		chi.URLParam(r, "resource_provider_uuid"),
		appplacement.UpdateTraits{
			ResourceProviderGeneration: request.ResourceProviderGeneration,
			Traits:                     request.Traits,
		},
	)
	if handlePlacementError(w, err, "trait update failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toTraitsDocument(traits))
}

func (h Handler) deleteTraits(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteTraits(chi.URLParam(r, "resource_provider_uuid"))
	if handlePlacementError(w, err, "trait delete failed") {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
