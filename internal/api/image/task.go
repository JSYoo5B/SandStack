package image

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listTasks(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, taskListResponse{
		Tasks: toTaskDocuments(h.service.ListTasks()),
		Next:  "",
	})
}

func (h Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var request createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	task := h.service.CreateTask(request.createTask())

	respond.JSON(w, http.StatusCreated, toTaskDocument(task))
}

func (h Handler) getTask(w http.ResponseWriter, r *http.Request) {
	task, err := h.service.GetTask(chi.URLParam(r, "task_id"))
	if errors.Is(err, appimage.ErrTaskNotFound) {
		respond.Error(w, http.StatusNotFound, "task not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "task lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, toTaskDocument(task))
}
