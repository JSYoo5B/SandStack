package image

import (
	"sync"

	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
)

type MemoryTaskRepository struct {
	mu    sync.RWMutex
	ids   []string
	tasks map[string]appimage.Task
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		ids:   []string{},
		tasks: map[string]appimage.Task{},
	}
}

func (r *MemoryTaskRepository) Create(task appimage.Task) appimage.Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, task.ID)
	r.tasks[task.ID] = task

	return task
}

func (r *MemoryTaskRepository) List() []appimage.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]appimage.Task, 0, len(r.ids))
	for _, id := range r.ids {
		tasks = append(tasks, r.tasks[id])
	}

	return tasks
}

func (r *MemoryTaskRepository) Get(id string) (appimage.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return appimage.Task{}, appimage.ErrTaskNotFound
	}

	return task, nil
}

func (r *MemoryTaskRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.tasks = map[string]appimage.Task{}
}
