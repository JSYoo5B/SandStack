package volume

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listAttachments(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, attachmentListResponse{
		Attachments: toAttachmentDocuments(h.service.ListAttachments()),
	})
}

func (h Handler) createAttachment(w http.ResponseWriter, r *http.Request) {
	var request createAttachmentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	attachment := h.service.CreateAttachment(request.createAttachment())

	respond.JSON(w, http.StatusAccepted, attachmentResponse{
		Attachment: toAttachmentDocument(attachment),
	})
}

func (h Handler) getAttachment(w http.ResponseWriter, r *http.Request) {
	attachment, err := h.service.GetAttachment(
		chi.URLParam(r, "attachment_id"),
	)
	if errors.Is(err, appvolume.ErrAttachmentNotFound) {
		respond.Error(w, http.StatusNotFound, "attachment not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "attachment lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, attachmentResponse{
		Attachment: toAttachmentDocument(attachment),
	})
}

func (h Handler) updateAttachment(w http.ResponseWriter, r *http.Request) {
	var request updateAttachmentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	attachment, err := h.service.UpdateAttachment(
		chi.URLParam(r, "attachment_id"),
		request.updateAttachment(),
	)
	if errors.Is(err, appvolume.ErrAttachmentNotFound) {
		respond.Error(w, http.StatusNotFound, "attachment not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "attachment update failed")
		return
	}

	respond.JSON(w, http.StatusOK, attachmentResponse{
		Attachment: toAttachmentDocument(attachment),
	})
}

func (h Handler) deleteAttachment(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteAttachment(chi.URLParam(r, "attachment_id"))
	if errors.Is(err, appvolume.ErrAttachmentNotFound) {
		respond.Error(w, http.StatusNotFound, "attachment not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "attachment delete failed")
		return
	}

	respond.JSON(w, http.StatusOK, map[string]any{})
}

func (h Handler) actionAttachment(w http.ResponseWriter, r *http.Request) {
	var request attachmentActionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	err := h.service.CompleteAttachment(chi.URLParam(r, "attachment_id"))
	if errors.Is(err, appvolume.ErrAttachmentNotFound) {
		respond.Error(w, http.StatusNotFound, "attachment not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "attachment action failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
