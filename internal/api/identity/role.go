package identity

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listRoles(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, rolesResponse{
		Roles: h.service.Roles(),
		Links: identityLinks{
			Self: h.baseURL(r) + "/identity/v3/roles",
		},
	})
}

func (h Handler) getRole(w http.ResponseWriter, r *http.Request) {
	role, ok := h.service.RoleByID(chi.URLParam(r, "role_id"))
	if !ok {
		respond.Error(w, http.StatusNotFound, "role not found")
		return
	}

	respond.JSON(w, http.StatusOK, roleResponse{
		Role: role,
	})
}
