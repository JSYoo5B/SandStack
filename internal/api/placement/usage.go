package placement

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/go-chi/chi/v5"
)

func (h Handler) getUsages(w http.ResponseWriter, r *http.Request) {
	usages, err := h.service.GetUsages(
		chi.URLParam(r, "resource_provider_uuid"),
	)
	if handlePlacementError(w, err, "usage lookup failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toUsagesDocument(usages))
}
