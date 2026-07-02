package placement

import (
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listTraits(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, traitListResponse{
		Traits: h.service.ListTraits(r.URL.Query().Get("name")),
	})
}

func (h Handler) getTrait(w http.ResponseWriter, r *http.Request) {
	_, err := h.service.GetTrait(chi.URLParam(r, "trait_name"))
	if errors.Is(err, appplacement.ErrTraitNotFound) {
		respond.Error(w, http.StatusNotFound, "trait not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "trait lookup failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) createTrait(w http.ResponseWriter, r *http.Request) {
	created := h.service.CreateTrait(chi.URLParam(r, "trait_name"))
	if !created {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Handler) deleteTrait(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteTrait(chi.URLParam(r, "trait_name"))
	if errors.Is(err, appplacement.ErrTraitNotFound) {
		respond.Error(w, http.StatusNotFound, "trait not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "trait delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
