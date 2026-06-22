package admin

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listRequests(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, requestListResponse{
		Requests: toRequestDocuments(h.requests.List()),
	})
}
