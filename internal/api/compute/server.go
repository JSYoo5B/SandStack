package compute

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listServers(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, serverListResponse{
		Servers: toServerDocuments(h.service.ListServers()),
	})
}
