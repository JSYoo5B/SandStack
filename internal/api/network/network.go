package network

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listNetworks(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, networkListResponse{
		Networks: []networkDocument{},
	})
}
