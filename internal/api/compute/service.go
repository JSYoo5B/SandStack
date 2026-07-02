package compute

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listComputeServices(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, computeServiceListResponse{
		Services: toComputeServiceDocuments(h.service.ListComputeServices()),
	})
}
