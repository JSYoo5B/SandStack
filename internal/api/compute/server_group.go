package compute

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listServerGroups(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, serverGroupListResponse{
		ServerGroups: toServerGroupDocuments(h.service.ListServerGroups()),
	})
}

func (h Handler) createServerGroup(w http.ResponseWriter, r *http.Request) {
	var request createServerGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid server group request")
		return
	}

	group := h.service.CreateServerGroup(request.createServerGroup())

	respond.JSON(w, http.StatusOK, serverGroupResponse{
		ServerGroup: toServerGroupDocument(group),
	})
}

func (h Handler) getServerGroup(w http.ResponseWriter, r *http.Request) {
	group, err := h.service.GetServerGroup(chi.URLParam(r, "group_id"))
	if errors.Is(err, appcompute.ErrServerGroupNotFound) {
		respond.Error(w, http.StatusNotFound, "server group not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "server group lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, serverGroupResponse{
		ServerGroup: toServerGroupDocument(group),
	})
}

func (h Handler) deleteServerGroup(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteServerGroup(chi.URLParam(r, "group_id"))
	if errors.Is(err, appcompute.ErrServerGroupNotFound) {
		respond.Error(w, http.StatusNotFound, "server group not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "server group delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
