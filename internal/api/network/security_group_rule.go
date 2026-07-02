package network

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listSecurityGroupRules(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, securityGroupRuleListResponse{
		SecurityGroupRules: toSecurityGroupRuleDocuments(
			h.service.ListSecurityGroupRules(),
		),
	})
}

func (h Handler) createSecurityGroupRule(w http.ResponseWriter, r *http.Request) {
	var request createSecurityGroupRuleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	rule, err := h.service.CreateSecurityGroupRule(
		request.createSecurityGroupRule(),
	)
	if errors.Is(err, appnetwork.ErrSecurityGroupNotFound) {
		respond.Error(w, http.StatusNotFound, "security group not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "security group rule create failed")
		return
	}

	respond.JSON(w, http.StatusCreated, securityGroupRuleResponse{
		SecurityGroupRule: toSecurityGroupRuleDocument(rule),
	})
}

func (h Handler) getSecurityGroupRule(w http.ResponseWriter, r *http.Request) {
	rule, err := h.service.GetSecurityGroupRule(
		chi.URLParam(r, "security_group_rule_id"),
	)
	if errors.Is(err, appnetwork.ErrSecurityGroupRuleNotFound) {
		respond.Error(w, http.StatusNotFound, "security group rule not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "security group rule lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, securityGroupRuleResponse{
		SecurityGroupRule: toSecurityGroupRuleDocument(rule),
	})
}

func (h Handler) deleteSecurityGroupRule(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteSecurityGroupRule(
		chi.URLParam(r, "security_group_rule_id"),
	)
	if errors.Is(err, appnetwork.ErrSecurityGroupRuleNotFound) {
		respond.Error(w, http.StatusNotFound, "security group rule not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "security group rule delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
