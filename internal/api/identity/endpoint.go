package identity

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listEndpoints(w http.ResponseWriter, r *http.Request) {
	baseURL := h.baseURL(r)
	respond.JSON(w, http.StatusOK, endpointsResponse{
		Endpoints: h.service.Endpoints(baseURL),
		Links: identityLinks{
			Self: baseURL + "/identity/v3/endpoints",
		},
	})
}

func (h Handler) getEndpoint(w http.ResponseWriter, r *http.Request) {
	baseURL := h.baseURL(r)
	endpoint, ok := h.service.EndpointByID(baseURL, chi.URLParam(r, "endpoint_id"))
	if !ok {
		respond.Error(w, http.StatusNotFound, "endpoint not found")
		return
	}

	respond.JSON(w, http.StatusOK, endpointResponse{
		Endpoint: endpoint,
	})
}
