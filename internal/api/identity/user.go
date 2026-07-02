package identity

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	baseURL := h.baseURL(r)
	users := h.service.Users()
	responseUsers := make([]user, 0, len(users))
	for _, currentUser := range users {
		responseUsers = append(responseUsers, toUser(currentUser, baseURL))
	}

	respond.JSON(w, http.StatusOK, usersResponse{
		Users: responseUsers,
		Links: identityLinks{
			Self: baseURL + "/identity/v3/users",
		},
	})
}

func (h Handler) getUser(w http.ResponseWriter, r *http.Request) {
	currentUser, ok := h.service.UserByID(chi.URLParam(r, "user_id"))
	if !ok {
		respond.Error(w, http.StatusNotFound, "user not found")
		return
	}

	respond.JSON(w, http.StatusOK, userResponse{
		User: toUser(currentUser, h.baseURL(r)),
	})
}
