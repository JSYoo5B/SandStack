package network

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listSubnets(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, subnetListResponse{
		Subnets: toSubnetDocuments(h.service.ListSubnets()),
	})
}

func (h Handler) createSubnet(w http.ResponseWriter, r *http.Request) {
	var request createSubnetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	subnet := h.service.CreateSubnet(request.createSubnet())
	respond.JSON(w, http.StatusCreated, subnetResponse{
		Subnet: toSubnetDocument(subnet),
	})
}

func (h Handler) getSubnet(w http.ResponseWriter, r *http.Request) {
	subnet, err := h.service.GetSubnet(chi.URLParam(r, "subnet_id"))
	if errors.Is(err, appnetwork.ErrSubnetNotFound) {
		respond.Error(w, http.StatusNotFound, "subnet not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "subnet lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, subnetResponse{
		Subnet: toSubnetDocument(subnet),
	})
}

func (h Handler) deleteSubnet(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteSubnet(chi.URLParam(r, "subnet_id"))
	if errors.Is(err, appnetwork.ErrSubnetNotFound) {
		respond.Error(w, http.StatusNotFound, "subnet not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "subnet delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
