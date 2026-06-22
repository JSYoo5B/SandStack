package network

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/go-chi/chi/v5"
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

func (h Handler) getPort(w http.ResponseWriter, r *http.Request) {
	port, err := h.service.GetPort(chi.URLParam(r, "port_id"))
	if errors.Is(err, appnetwork.ErrPortNotFound) {
		respond.Error(w, http.StatusNotFound, "port not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "port lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, portResponse{
		Port: toPortDocument(port),
	})
}

func (h Handler) deletePort(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeletePort(chi.URLParam(r, "port_id"))
	if errors.Is(err, appnetwork.ErrPortNotFound) {
		respond.Error(w, http.StatusNotFound, "port not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "port delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
