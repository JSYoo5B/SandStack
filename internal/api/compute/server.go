package compute

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
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

func (h Handler) getServer(w http.ResponseWriter, r *http.Request) {
	server, err := h.service.GetServer(chi.URLParam(r, "server_id"))
	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "server lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, serverResponse{
		Server: toServerDocument(server),
	})
}

func (h Handler) deleteServer(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteServer(chi.URLParam(r, "server_id"))
	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "server delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
