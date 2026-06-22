package network

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listNetworks(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, networkListResponse{
		Networks: toNetworkDocuments(h.service.List()),
	})
}

func (h Handler) createNetwork(w http.ResponseWriter, r *http.Request) {
	var request createNetworkRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	network := h.service.Create(request.createNetwork())
	respond.JSON(w, http.StatusCreated, networkResponse{
		Network: toNetworkDocument(network),
	})
}

func (h Handler) getNetwork(w http.ResponseWriter, r *http.Request) {
	network, err := h.service.Get(chi.URLParam(r, "network_id"))
	if errors.Is(err, appnetwork.ErrNetworkNotFound) {
		respond.Error(w, http.StatusNotFound, "network not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "network lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, networkResponse{
		Network: toNetworkDocument(network),
	})
}

func (h Handler) deleteNetwork(w http.ResponseWriter, r *http.Request) {
	err := h.service.Delete(chi.URLParam(r, "network_id"))
	if errors.Is(err, appnetwork.ErrNetworkNotFound) {
		respond.Error(w, http.StatusNotFound, "network not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "network delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
