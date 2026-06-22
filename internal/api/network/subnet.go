package network

import (
	"encoding/json"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
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
