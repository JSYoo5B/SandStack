package compute

import (
	"encoding/json"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listServers(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, serverListResponse{
		Servers: toServerDocuments(h.service.ListServers()),
	})
}

func (h Handler) createServer(w http.ResponseWriter, r *http.Request) {
	var request createServerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	server := h.service.CreateServer(request.createServer())
	respond.JSON(w, http.StatusAccepted, serverResponse{
		Server: toServerDocument(server),
	})
}
