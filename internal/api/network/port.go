package network

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listPorts(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, portListResponse{
		Ports: toPortDocuments(h.service.ListPorts()),
	})
}
