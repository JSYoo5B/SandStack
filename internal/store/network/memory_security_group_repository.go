package network

import (
	"sync"

	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
)

type MemorySecurityGroupRepository struct {
	mu             sync.RWMutex
	ids            []string
	securityGroups map[string]appnetwork.SecurityGroup
}

func NewMemorySecurityGroupRepository() *MemorySecurityGroupRepository {
	return &MemorySecurityGroupRepository{
		ids:            []string{},
		securityGroups: map[string]appnetwork.SecurityGroup{},
	}
}

func (r *MemorySecurityGroupRepository) Create(
	securityGroup appnetwork.SecurityGroup,
) appnetwork.SecurityGroup {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, securityGroup.ID)
	r.securityGroups[securityGroup.ID] = securityGroup

	return securityGroup
}

func (r *MemorySecurityGroupRepository) List() []appnetwork.SecurityGroup {
	r.mu.RLock()
	defer r.mu.RUnlock()

	securityGroups := make([]appnetwork.SecurityGroup, 0, len(r.ids))
	for _, id := range r.ids {
		securityGroups = append(securityGroups, r.securityGroups[id])
	}

	return securityGroups
}

func (r *MemorySecurityGroupRepository) Get(
	id string,
) (appnetwork.SecurityGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	securityGroup, ok := r.securityGroups[id]
	if !ok {
		return appnetwork.SecurityGroup{}, appnetwork.ErrSecurityGroupNotFound
	}

	return securityGroup, nil
}

func (r *MemorySecurityGroupRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.securityGroups[id]; !ok {
		return appnetwork.ErrSecurityGroupNotFound
	}

	delete(r.securityGroups, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemorySecurityGroupRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.securityGroups = map[string]appnetwork.SecurityGroup{}
}
