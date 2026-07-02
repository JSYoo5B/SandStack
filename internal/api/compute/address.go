package compute

import (
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listServerAddresses(w http.ResponseWriter, r *http.Request) {
	addresses, err := h.service.ServerAddresses(chi.URLParam(r, "server_id"))
	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "server address lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, serverAddressListResponse{
		Addresses: toServerAddressDocumentsByNetwork(addresses),
	})
}

func (h Handler) listServerAddressesByNetwork(w http.ResponseWriter, r *http.Request) {
	addresses, err := h.service.ServerAddressesByNetwork(
		chi.URLParam(r, "server_id"),
		chi.URLParam(r, "network"),
	)
	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "server address lookup failed")
		return
	}

	network := chi.URLParam(r, "network")
	respond.JSON(w, http.StatusOK, map[string][]serverAddressDocument{
		network: toServerAddressDocuments(addresses),
	})
}
