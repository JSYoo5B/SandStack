package network

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listFloatingIPs(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, floatingIPListResponse{
		FloatingIPs: toFloatingIPDocuments(h.service.ListFloatingIPs()),
	})
}

func (h Handler) createFloatingIP(w http.ResponseWriter, r *http.Request) {
	var request createFloatingIPRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	floatingIP := h.service.CreateFloatingIP(request.createFloatingIP())
	respond.JSON(w, http.StatusCreated, floatingIPResponse{
		FloatingIP: toFloatingIPDocument(floatingIP),
	})
}

func (h Handler) getFloatingIP(w http.ResponseWriter, r *http.Request) {
	floatingIP, err := h.service.GetFloatingIP(chi.URLParam(r, "floating_ip_id"))
	if errors.Is(err, appnetwork.ErrFloatingIPNotFound) {
		respond.Error(w, http.StatusNotFound, "floating IP not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "floating IP lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, floatingIPResponse{
		FloatingIP: toFloatingIPDocument(floatingIP),
	})
}

func (h Handler) deleteFloatingIP(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteFloatingIP(chi.URLParam(r, "floating_ip_id"))
	if errors.Is(err, appnetwork.ErrFloatingIPNotFound) {
		respond.Error(w, http.StatusNotFound, "floating IP not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "floating IP delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
