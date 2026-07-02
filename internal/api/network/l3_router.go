package network

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listRouters(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, routerListResponse{
		Routers: toRouterDocuments(h.service.ListRouters()),
	})
}

func (h Handler) createRouter(w http.ResponseWriter, r *http.Request) {
	var request createRouterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	router := h.service.CreateRouter(request.createRouter())
	respond.JSON(w, http.StatusCreated, routerResponse{
		Router: toRouterDocument(router),
	})
}

func (h Handler) getRouter(w http.ResponseWriter, r *http.Request) {
	router, err := h.service.GetRouter(chi.URLParam(r, "router_id"))
	if errors.Is(err, appnetwork.ErrRouterNotFound) {
		respond.Error(w, http.StatusNotFound, "router not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "router lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, routerResponse{
		Router: toRouterDocument(router),
	})
}

func (h Handler) deleteRouter(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteRouter(chi.URLParam(r, "router_id"))
	if errors.Is(err, appnetwork.ErrRouterNotFound) {
		respond.Error(w, http.StatusNotFound, "router not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "router delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
