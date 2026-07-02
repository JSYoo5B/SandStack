package identity

import (
	"sync"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"
)

type MemoryRoleRepository struct {
	mu    sync.RWMutex
	ids   []string
	roles map[string]roles.Role
}

func NewMemoryRoleRepository() *MemoryRoleRepository {
	return &MemoryRoleRepository{
		roles: map[string]roles.Role{},
	}
}

func (r *MemoryRoleRepository) Save(role roles.Role) roles.Role {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.roles[role.ID]; !ok {
		r.ids = append(r.ids, role.ID)
	}
	r.roles[role.ID] = role

	return role
}

func (r *MemoryRoleRepository) List() []roles.Role {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]roles.Role, 0, len(r.ids))
	for _, id := range r.ids {
		result = append(result, r.roles[id])
	}

	return result
}

func (r *MemoryRoleRepository) Get(id string) (roles.Role, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	role, ok := r.roles[id]
	if !ok {
		return roles.Role{}, appidentity.ErrRoleNotFound
	}

	return role, nil
}

func (r *MemoryRoleRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = nil
	r.roles = map[string]roles.Role{}
}
