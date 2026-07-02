package network

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/go-chi/chi/v5"
)

func (h Handler) addRouterInterface(w http.ResponseWriter, r *http.Request) {
	var request routerInterfaceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	routerInterface, err := h.service.AddRouterInterface(
		chi.URLParam(r, "router_id"),
		request.routerInterfaceRequest(),
	)
	if errors.Is(err, appnetwork.ErrRouterNotFound) {
		respond.Error(w, http.StatusNotFound, "router not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "router interface add failed")
		return
	}

	respond.JSON(w, http.StatusOK, toRouterInterfaceDocument(routerInterface))
}

func (h Handler) removeRouterInterface(w http.ResponseWriter, r *http.Request) {
	var request routerInterfaceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	routerInterface, err := h.service.RemoveRouterInterface(
		chi.URLParam(r, "router_id"),
		request.routerInterfaceRequest(),
	)
	if errors.Is(err, appnetwork.ErrRouterNotFound) {
		respond.Error(w, http.StatusNotFound, "router not found")
		return
	}
	if errors.Is(err, appnetwork.ErrRouterInterfaceNotFound) {
		respond.Error(w, http.StatusNotFound, "router interface not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "router interface remove failed")
		return
	}

	respond.JSON(w, http.StatusOK, toRouterInterfaceDocument(routerInterface))
}
