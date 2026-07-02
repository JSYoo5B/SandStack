package network

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listSecurityGroups(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, securityGroupListResponse{
		SecurityGroups: toSecurityGroupDocuments(
			h.service.ListSecurityGroups(),
		),
	})
}

func (h Handler) createSecurityGroup(w http.ResponseWriter, r *http.Request) {
	var request createSecurityGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	securityGroup := h.service.CreateSecurityGroup(
		request.createSecurityGroup(),
	)
	respond.JSON(w, http.StatusCreated, securityGroupResponse{
		SecurityGroup: toSecurityGroupDocument(securityGroup),
	})
}

func (h Handler) getSecurityGroup(w http.ResponseWriter, r *http.Request) {
	securityGroup, err := h.service.GetSecurityGroup(
		chi.URLParam(r, "security_group_id"),
	)
	if errors.Is(err, appnetwork.ErrSecurityGroupNotFound) {
		respond.Error(w, http.StatusNotFound, "security group not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "security group lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, securityGroupResponse{
		SecurityGroup: toSecurityGroupDocument(securityGroup),
	})
}

func (h Handler) deleteSecurityGroup(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteSecurityGroup(
		chi.URLParam(r, "security_group_id"),
	)
	if errors.Is(err, appnetwork.ErrSecurityGroupNotFound) {
		respond.Error(w, http.StatusNotFound, "security group not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "security group delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
