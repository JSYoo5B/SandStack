package placement

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listResourceClasses(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, resourceClassListResponse{
		ResourceClasses: toResourceClassDocuments(
			h.service.ListResourceClasses(),
		),
	})
}

func (h Handler) createResourceClass(w http.ResponseWriter, r *http.Request) {
	var request createResourceClassRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	h.service.CreateResourceClass(request.Name)
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) getResourceClass(w http.ResponseWriter, r *http.Request) {
	class, err := h.service.GetResourceClass(
		chi.URLParam(r, "resource_class"),
	)
	if errors.Is(err, appplacement.ErrResourceClassNotFound) {
		respond.Error(w, http.StatusNotFound, "resource class not found")
		return
	}
	if err != nil {
		respond.Error(
			w,
			http.StatusInternalServerError,
			"resource class lookup failed",
		)
		return
	}

	respond.JSON(w, http.StatusOK, toResourceClassDocument(class))
}

func (h Handler) updateResourceClass(w http.ResponseWriter, r *http.Request) {
	_, err := h.service.GetResourceClass(
		chi.URLParam(r, "resource_class"),
	)

	h.service.EnsureResourceClass(chi.URLParam(r, "resource_class"))
	if err == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Handler) deleteResourceClass(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteResourceClass(chi.URLParam(r, "resource_class"))
	if errors.Is(err, appplacement.ErrResourceClassNotFound) {
		respond.Error(w, http.StatusNotFound, "resource class not found")
		return
	}
	if err != nil {
		respond.Error(
			w,
			http.StatusInternalServerError,
			"resource class delete failed",
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
