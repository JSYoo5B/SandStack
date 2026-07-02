package volume

import (
	"encoding/json"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/go-chi/chi/v5"
)

func (h Handler) getQuotaSet(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "quota_project_id")
	if r.URL.Query().Get("usage") == "true" {
		respond.JSON(w, http.StatusOK, quotaUsageSetResponse{
			QuotaSet: toQuotaUsageSetDocument(
				h.service.GetQuotaUsageSet(projectID),
			),
		})
		return
	}

	respond.JSON(w, http.StatusOK, quotaSetResponse{
		QuotaSet: toQuotaSetDocument(h.service.GetQuotaSet(projectID)),
	})
}

func (h Handler) getDefaultQuotaSet(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, quotaSetResponse{
		QuotaSet: toQuotaSetDocument(
			h.service.GetDefaultQuotaSet(
				chi.URLParam(r, "quota_project_id"),
			),
		),
	})
}

func (h Handler) updateQuotaSet(w http.ResponseWriter, r *http.Request) {
	var request updateQuotaSetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid quota request")
		return
	}

	quotaSet := h.service.UpdateQuotaSet(
		chi.URLParam(r, "quota_project_id"),
		request.updateQuotaSet(),
	)

	respond.JSON(w, http.StatusOK, quotaSetResponse{
		QuotaSet: toQuotaSetDocument(quotaSet),
	})
}

func (h Handler) deleteQuotaSet(w http.ResponseWriter, r *http.Request) {
	if err := h.service.ResetQuotaSet(
		chi.URLParam(r, "quota_project_id"),
	); err != nil {
		respond.Error(w, http.StatusInternalServerError, "quota reset failed")
		return
	}

	respond.JSON(w, http.StatusOK, map[string]any{})
}
