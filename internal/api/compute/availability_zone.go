package compute

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func (h Handler) listAvailabilityZones(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, availabilityZoneListResponse{
		AvailabilityZoneInfo: toAvailabilityZoneDocuments(
			h.service.ListAvailabilityZones(),
		),
	})
}
