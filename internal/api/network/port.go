package network

import (
	"encoding/json"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listPorts(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, portListResponse{
		Ports: toPortDocuments(h.service.ListPorts()),
	})
}

func (h Handler) createPort(w http.ResponseWriter, r *http.Request) {
	var request createPortRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	port := h.service.CreatePort(request.createPort())
	respond.JSON(w, http.StatusCreated, portResponse{
		Port: toPortDocument(port),
	})
}
