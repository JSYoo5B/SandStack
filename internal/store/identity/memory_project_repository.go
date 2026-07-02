package identity

import (
	"sync"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"
)

type MemoryProjectRepository struct {
	mu       sync.RWMutex
	ids      []string
	projects map[string]projects.Project
}

func NewMemoryProjectRepository() *MemoryProjectRepository {
	return &MemoryProjectRepository{
		projects: map[string]projects.Project{},
	}
}

func (r *MemoryProjectRepository) Save(
	project projects.Project,
) projects.Project {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.projects[project.ID]; !ok {
		r.ids = append(r.ids, project.ID)
	}
	r.projects[project.ID] = project

	return project
}

func (r *MemoryProjectRepository) List() []projects.Project {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]projects.Project, 0, len(r.ids))
	for _, id := range r.ids {
		result = append(result, r.projects[id])
	}

	return result
}

func (r *MemoryProjectRepository) Get(
	id string,
) (projects.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	project, ok := r.projects[id]
	if !ok {
		return projects.Project{}, appidentity.ErrProjectNotFound
	}

	return project, nil
}

func (r *MemoryProjectRepository) FindByName(
	name string,
) (projects.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, id := range r.ids {
		project := r.projects[id]
		if project.Name == name {
			return project, nil
		}
	}

	return projects.Project{}, appidentity.ErrProjectNotFound
}

func (r *MemoryProjectRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = nil
	r.projects = map[string]projects.Project{}
}
