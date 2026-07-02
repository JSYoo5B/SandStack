package image

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listMembers(w http.ResponseWriter, r *http.Request) {
	members, err := h.service.ListMembers(chi.URLParam(r, "image_id"))
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image member lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, memberListResponse{
		Members: toMemberDocuments(members),
	})
}

func (h Handler) createMember(w http.ResponseWriter, r *http.Request) {
	var request createMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	member, err := h.service.CreateMember(
		chi.URLParam(r, "image_id"),
		request.Member,
	)
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image member create failed")
		return
	}

	respond.JSON(w, http.StatusOK, toMemberDocument(member))
}

func (h Handler) getMember(w http.ResponseWriter, r *http.Request) {
	member, err := h.service.GetMember(
		chi.URLParam(r, "image_id"),
		chi.URLParam(r, "member_id"),
	)
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image member not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image member lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, toMemberDocument(member))
}

func (h Handler) updateMember(w http.ResponseWriter, r *http.Request) {
	var request updateMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	member, err := h.service.UpdateMember(
		chi.URLParam(r, "image_id"),
		chi.URLParam(r, "member_id"),
		request.Status,
	)
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image member not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image member update failed")
		return
	}

	respond.JSON(w, http.StatusOK, toMemberDocument(member))
}

func (h Handler) deleteMember(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteMember(
		chi.URLParam(r, "image_id"),
		chi.URLParam(r, "member_id"),
	)
	if errors.Is(err, appimage.ErrImageNotFound) {
		respond.Error(w, http.StatusNotFound, "image member not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "image member delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
