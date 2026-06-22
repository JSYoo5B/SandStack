package network

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listSubnets(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, subnetListResponse{
		Subnets: toSubnetDocuments(h.service.ListSubnets()),
	})
}
