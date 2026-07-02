package identity

import (
	"sync"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
)

type MemoryServiceRepository struct {
	mu       sync.RWMutex
	ids      []string
	services map[string]appidentity.ServiceDefinition
}

func NewMemoryServiceRepository() *MemoryServiceRepository {
	return &MemoryServiceRepository{
		services: map[string]appidentity.ServiceDefinition{},
	}
}

func (r *MemoryServiceRepository) Save(
	service appidentity.ServiceDefinition,
) appidentity.ServiceDefinition {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.services[service.ID]; !ok {
		r.ids = append(r.ids, service.ID)
	}
	r.services[service.ID] = service

	return service
}

func (r *MemoryServiceRepository) List() []appidentity.ServiceDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	services := make([]appidentity.ServiceDefinition, 0, len(r.ids))
	for _, id := range r.ids {
		services = append(services, r.services[id])
	}

	return services
}

func (r *MemoryServiceRepository) Get(
	id string,
) (appidentity.ServiceDefinition, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	service, ok := r.services[id]
	if !ok {
		return appidentity.ServiceDefinition{}, appidentity.ErrServiceNotFound
	}

	return service, nil
}

func (r *MemoryServiceRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = nil
	r.services = map[string]appidentity.ServiceDefinition{}
}
