package compute

import (
	"sync"

	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
)

type MemoryServerGroupRepository struct {
	mu     sync.RWMutex
	ids    []string
	groups map[string]appcompute.ServerGroup
}

func NewMemoryServerGroupRepository() *MemoryServerGroupRepository {
	return &MemoryServerGroupRepository{
		groups: map[string]appcompute.ServerGroup{},
	}
}

func (r *MemoryServerGroupRepository) Create(
	group appcompute.ServerGroup,
) appcompute.ServerGroup {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, group.ID)
	r.groups[group.ID] = group

	return group
}

func (r *MemoryServerGroupRepository) List() []appcompute.ServerGroup {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groups := make([]appcompute.ServerGroup, 0, len(r.ids))
	for _, id := range r.ids {
		groups = append(groups, r.groups[id])
	}

	return groups
}

func (r *MemoryServerGroupRepository) Get(
	id string,
) (appcompute.ServerGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	group, ok := r.groups[id]
	if !ok {
		return appcompute.ServerGroup{}, appcompute.ErrServerGroupNotFound
	}

	return group, nil
}

func (r *MemoryServerGroupRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.groups[id]; !ok {
		return appcompute.ErrServerGroupNotFound
	}
	delete(r.groups, id)
	for index, existingID := range r.ids {
		if existingID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryServerGroupRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = nil
	r.groups = map[string]appcompute.ServerGroup{}
}
