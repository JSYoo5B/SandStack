package image

import "time"

func (s *Service) CreateTask(input CreateTask) Task {
	now := s.clock.Now().UTC()
	task := Task{
		ID:        "task-" + s.idGen.Hex(16),
		Type:      input.Type,
		Status:    "pending",
		Input:     input.Input,
		Result:    map[string]any{},
		Owner:     "admin",
		ExpiresAt: now.Add(24 * time.Hour).Format(time.RFC3339),
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}

	return s.taskRepository.Create(task)
}

func (s *Service) ListTasks() []Task {
	return s.taskRepository.List()
}

func (s *Service) GetTask(id string) (Task, error) {
	return s.taskRepository.Get(id)
}
