package placement

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listResourceProviders(w http.ResponseWriter, r *http.Request) {
	providers := h.service.ListResourceProviders()

	respond.JSON(w, http.StatusOK, resourceProviderListResponse{
		ResourceProviders: toResourceProviderDocuments(providers),
	})
}

func (h Handler) createResourceProvider(w http.ResponseWriter, r *http.Request) {
	var request createResourceProviderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	provider := h.service.CreateResourceProvider(
		appplacement.CreateResourceProvider{
			UUID:               request.UUID,
			Name:               request.Name,
			ParentProviderUUID: request.ParentProviderUUID,
		},
	)

	respond.JSON(w, http.StatusOK, toResourceProviderDocument(provider))
}

func (h Handler) getResourceProvider(w http.ResponseWriter, r *http.Request) {
	provider, err := h.service.GetResourceProvider(
		chi.URLParam(r, "resource_provider_uuid"),
	)
	if errors.Is(err, appplacement.ErrResourceProviderNotFound) {
		respond.Error(w, http.StatusNotFound, "resource provider not found")
		return
	}
	if err != nil {
		respond.Error(
			w,
			http.StatusInternalServerError,
			"resource provider lookup failed",
		)
		return
	}

	respond.JSON(w, http.StatusOK, toResourceProviderDocument(provider))
}

func (h Handler) deleteResourceProvider(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteResourceProvider(
		chi.URLParam(r, "resource_provider_uuid"),
	)
	if errors.Is(err, appplacement.ErrResourceProviderNotFound) {
		respond.Error(w, http.StatusNotFound, "resource provider not found")
		return
	}
	if err != nil {
		respond.Error(
			w,
			http.StatusInternalServerError,
			"resource provider delete failed",
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
