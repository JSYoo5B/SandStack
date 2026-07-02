package identity

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listServices(w http.ResponseWriter, r *http.Request) {
	baseURL := h.baseURL(r)
	respond.JSON(w, http.StatusOK, servicesResponse{
		Services: h.service.Services(baseURL),
		Links: identityLinks{
			Self: baseURL + "/identity/v3/services",
		},
	})
}

func (h Handler) getService(w http.ResponseWriter, r *http.Request) {
	baseURL := h.baseURL(r)
	service, ok := h.service.ServiceByID(baseURL, chi.URLParam(r, "service_id"))
	if !ok {
		respond.Error(w, http.StatusNotFound, "service not found")
		return
	}

	respond.JSON(w, http.StatusOK, serviceResponse{
		Service: service,
	})
}
