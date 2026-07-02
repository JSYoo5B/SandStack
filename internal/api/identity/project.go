package identity

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listProjects(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, projectsResponse{
		Projects: h.service.Projects(),
		Links: identityLinks{
			Self: h.baseURL(r) + "/identity/v3/projects",
		},
	})
}

func (h Handler) getProject(w http.ResponseWriter, r *http.Request) {
	project, ok := h.service.ProjectByID(chi.URLParam(r, "project_id"))
	if !ok {
		respond.Error(w, http.StatusNotFound, "project not found")
		return
	}

	respond.JSON(w, http.StatusOK, projectResponse{
		Project: project,
	})
}
