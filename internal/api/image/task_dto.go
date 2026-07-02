package image

import appimage "github.com/JSYoo5B/SandStack/internal/app/image"

type createTaskRequest struct {
	Type  string         `json:"type"`
	Input map[string]any `json:"input"`
}

type taskListResponse struct {
	Tasks []taskDocument `json:"tasks"`
	Next  string         `json:"next"`
}

type taskDocument struct {
	ID        string         `json:"id"`
	Type      string         `json:"type"`
	Status    string         `json:"status"`
	Input     map[string]any `json:"input"`
	Result    map[string]any `json:"result"`
	Owner     string         `json:"owner"`
	Message   string         `json:"message"`
	ExpiresAt string         `json:"expires_at"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	Self      string         `json:"self"`
	Schema    string         `json:"schema"`
}

func (r createTaskRequest) createTask() appimage.CreateTask {
	return appimage.CreateTask{
		Type:  r.Type,
		Input: r.Input,
	}
}

func toTaskDocuments(tasks []appimage.Task) []taskDocument {
	documents := make([]taskDocument, 0, len(tasks))
	for _, task := range tasks {
		documents = append(documents, toTaskDocument(task))
	}

	return documents
}

func toTaskDocument(task appimage.Task) taskDocument {
	return taskDocument{
		ID:        task.ID,
		Type:      task.Type,
		Status:    task.Status,
		Input:     task.Input,
		Result:    task.Result,
		Owner:     task.Owner,
		Message:   task.Message,
		ExpiresAt: task.ExpiresAt,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		Self:      "/v2/tasks/" + task.ID,
		Schema:    "/v2/schemas/task",
	}
}
