package compute

import (
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listFlavors(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, flavorListResponse{
		Flavors: toFlavorDocuments(h.service.ListFlavors()),
	})
}

func (h Handler) getFlavor(w http.ResponseWriter, r *http.Request) {
	flavor, err := h.service.GetFlavor(chi.URLParam(r, "flavor_id"))
	if errors.Is(err, appcompute.ErrFlavorNotFound) {
		respond.Error(w, http.StatusNotFound, "flavor not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "flavor lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, flavorResponse{
		Flavor: toFlavorDocument(flavor),
	})
}
