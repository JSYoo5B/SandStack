package placement

import (
	"encoding/json"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
	"github.com/go-chi/chi/v5"
)

func (h Handler) getAggregates(w http.ResponseWriter, r *http.Request) {
	aggregates, err := h.service.GetAggregates(
		chi.URLParam(r, "resource_provider_uuid"),
	)
	if handlePlacementError(w, err, "aggregate lookup failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toAggregatesDocument(aggregates))
}

func (h Handler) updateAggregates(w http.ResponseWriter, r *http.Request) {
	var request aggregatesDocument
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	aggregates, err := h.service.UpdateAggregates(
		chi.URLParam(r, "resource_provider_uuid"),
		appplacement.UpdateAggregates{
			ResourceProviderGeneration: request.ResourceProviderGeneration,
			Aggregates:                 request.Aggregates,
		},
	)
	if handlePlacementError(w, err, "aggregate update failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toAggregatesDocument(aggregates))
}
