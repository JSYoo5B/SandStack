package network

import (
	"encoding/json"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
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
