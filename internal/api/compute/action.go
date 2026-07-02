package compute

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
)

func (h Handler) actionServer(w http.ResponseWriter, r *http.Request) {
	var request serverActionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	var err error
	switch {
	case request.Has("os-start"):
		err = h.service.StartServer(chi.URLParam(r, "server_id"))
	case request.Has("os-stop"):
		err = h.service.StopServer(chi.URLParam(r, "server_id"))
	case request.Has("reboot"):
		err = h.service.RebootServer(chi.URLParam(r, "server_id"))
	default:
		respond.Error(w, http.StatusBadRequest, "unsupported server action")
		return
	}

	if errors.Is(err, appcompute.ErrServerNotFound) {
		respond.Error(w, http.StatusNotFound, "server not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "server action failed")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
