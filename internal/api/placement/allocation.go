package placement

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/go-chi/chi/v5"
)

func (h Handler) getAllocations(w http.ResponseWriter, r *http.Request) {
	allocations, err := h.service.GetAllocations(
		chi.URLParam(r, "resource_provider_uuid"),
	)
	if handlePlacementError(w, err, "allocation lookup failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toAllocationsDocument(allocations))
}
