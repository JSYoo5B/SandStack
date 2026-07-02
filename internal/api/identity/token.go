package identity

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
)

func (h Handler) createToken(w http.ResponseWriter, r *http.Request) {
	var request authRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	issuedToken, err := h.service.AuthenticatePassword(
		request.passwordAuth(),
		h.baseURL(r),
	)
	if errors.Is(err, appidentity.ErrInvalidCredentials) {
		respond.Error(w, http.StatusUnauthorized, "authentication failed")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "authentication failed")
		return
	}

	w.Header().Set("X-Subject-Token", issuedToken.ID)
	respond.JSON(w, http.StatusCreated, tokenResponse{
		Token: toIssuedToken(issuedToken),
	})
}

func (h Handler) getToken(w http.ResponseWriter, r *http.Request) {
	tokenID := r.Header.Get("X-Subject-Token")
	issuedToken, err := h.service.TokenByID(tokenID)
	if errors.Is(err, appidentity.ErrTokenNotFound) {
		respond.Error(w, http.StatusNotFound, "token not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "token lookup failed")
		return
	}

	w.Header().Set("X-Subject-Token", issuedToken.ID)
	respond.JSON(w, http.StatusOK, tokenResponse{
		Token: toIssuedToken(issuedToken),
	})
}

func (h Handler) validateToken(w http.ResponseWriter, r *http.Request) {
	tokenID := r.Header.Get("X-Subject-Token")
	issuedToken, err := h.service.TokenByID(tokenID)
	if errors.Is(err, appidentity.ErrTokenNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-Subject-Token", issuedToken.ID)
	w.WriteHeader(http.StatusOK)
}

func (h Handler) revokeToken(w http.ResponseWriter, r *http.Request) {
	err := h.service.RevokeToken(r.Header.Get("X-Subject-Token"))
	if errors.Is(err, appidentity.ErrTokenNotFound) {
		respond.Error(w, http.StatusNotFound, "token not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "token revoke failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
