package placement

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
	"github.com/go-chi/chi/v5"
)

func (h Handler) getInventories(w http.ResponseWriter, r *http.Request) {
	inventories, err := h.service.GetInventories(
		chi.URLParam(r, "resource_provider_uuid"),
	)
	if handlePlacementError(w, err, "inventory lookup failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toInventoriesDocument(inventories))
}

func (h Handler) updateInventories(w http.ResponseWriter, r *http.Request) {
	var request inventoriesDocument
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	inventories, err := h.service.UpdateInventories(
		chi.URLParam(r, "resource_provider_uuid"),
		appplacement.UpdateInventories{
			ResourceProviderGeneration: request.ResourceProviderGeneration,
			Inventories:                toAppInventories(request),
		},
	)
	if handlePlacementError(w, err, "inventory update failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toInventoriesDocument(inventories))
}

func (h Handler) deleteInventories(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteInventories(chi.URLParam(r, "resource_provider_uuid"))
	if handlePlacementError(w, err, "inventory delete failed") {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) getInventory(w http.ResponseWriter, r *http.Request) {
	inventory, err := h.service.GetInventory(
		chi.URLParam(r, "resource_provider_uuid"),
		chi.URLParam(r, "resource_class"),
	)
	if handlePlacementError(w, err, "inventory lookup failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toInventoryWithGenerationDocument(inventory))
}

func (h Handler) updateInventory(w http.ResponseWriter, r *http.Request) {
	var request inventoryDocument
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	inventory, err := h.service.UpdateInventory(
		chi.URLParam(r, "resource_provider_uuid"),
		chi.URLParam(r, "resource_class"),
		appplacement.UpdateInventory{
			ResourceProviderGeneration: request.ResourceProviderGeneration,
			Inventory: toAppInventory(
				chi.URLParam(r, "resource_class"),
				request,
			),
		},
	)
	if handlePlacementError(w, err, "inventory update failed") {
		return
	}

	respond.JSON(w, http.StatusOK, toInventoryWithGenerationDocument(inventory))
}

func (h Handler) deleteInventory(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteInventory(
		chi.URLParam(r, "resource_provider_uuid"),
		chi.URLParam(r, "resource_class"),
	)
	if handlePlacementError(w, err, "inventory delete failed") {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handlePlacementError(
	w http.ResponseWriter,
	err error,
	internalMessage string,
) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, appplacement.ErrResourceProviderNotFound) {
		respond.Error(w, http.StatusNotFound, "resource provider not found")
		return true
	}
	if errors.Is(err, appplacement.ErrInventoryNotFound) {
		respond.Error(w, http.StatusNotFound, "inventory not found")
		return true
	}

	respond.Error(w, http.StatusInternalServerError, internalMessage)
	return true
}
