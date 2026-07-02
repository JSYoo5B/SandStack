package identity

import (
	"sync"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
)

type MemoryEndpointRepository struct {
	mu        sync.RWMutex
	ids       []string
	endpoints map[string]appidentity.EndpointDefinition
}

func NewMemoryEndpointRepository() *MemoryEndpointRepository {
	return &MemoryEndpointRepository{
		endpoints: map[string]appidentity.EndpointDefinition{},
	}
}

func (r *MemoryEndpointRepository) Save(
	endpoint appidentity.EndpointDefinition,
) appidentity.EndpointDefinition {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.endpoints[endpoint.ID]; !ok {
		r.ids = append(r.ids, endpoint.ID)
	}
	r.endpoints[endpoint.ID] = endpoint

	return endpoint
}

func (r *MemoryEndpointRepository) List() []appidentity.EndpointDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	endpoints := make([]appidentity.EndpointDefinition, 0, len(r.ids))
	for _, id := range r.ids {
		endpoints = append(endpoints, r.endpoints[id])
	}

	return endpoints
}

func (r *MemoryEndpointRepository) Get(
	id string,
) (appidentity.EndpointDefinition, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	endpoint, ok := r.endpoints[id]
	if !ok {
		return appidentity.EndpointDefinition{}, appidentity.ErrEndpointNotFound
	}

	return endpoint, nil
}

func (r *MemoryEndpointRepository) ListByServiceID(
	serviceID string,
) []appidentity.EndpointDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	endpoints := []appidentity.EndpointDefinition{}
	for _, id := range r.ids {
		endpoint := r.endpoints[id]
		if endpoint.ServiceID == serviceID {
			endpoints = append(endpoints, endpoint)
		}
	}

	return endpoints
}

func (r *MemoryEndpointRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = nil
	r.endpoints = map[string]appidentity.EndpointDefinition{}
}
