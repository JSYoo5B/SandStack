package compute

import (
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listHypervisors(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, hypervisorListResponse{
		Hypervisors: toHypervisorDocuments(h.service.ListHypervisors()),
	})
}

func (h Handler) getHypervisorStatistics(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, hypervisorStatisticsResponse{
		HypervisorStatistics: toHypervisorStatisticsDocument(
			h.service.ListHypervisors(),
		),
	})
}

func (h Handler) getHypervisor(w http.ResponseWriter, r *http.Request) {
	hypervisor, err := h.service.GetHypervisor(chi.URLParam(r, "hypervisor_id"))
	if errors.Is(err, appcompute.ErrHypervisorNotFound) {
		respond.Error(w, http.StatusNotFound, "hypervisor not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "hypervisor lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, hypervisorResponse{
		Hypervisor: toHypervisorDocument(hypervisor),
	})
}
