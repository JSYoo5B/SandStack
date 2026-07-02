package placement

import (
	"sync"

	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
)

type MemoryResourceProviderRepository struct {
	mu        sync.RWMutex
	uuids     []string
	providers map[string]appplacement.ResourceProvider
}

func NewMemoryResourceProviderRepository() *MemoryResourceProviderRepository {
	return &MemoryResourceProviderRepository{
		uuids:     []string{},
		providers: map[string]appplacement.ResourceProvider{},
	}
}

func (r *MemoryResourceProviderRepository) Create(
	provider appplacement.ResourceProvider,
) appplacement.ResourceProvider {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.providers[provider.UUID]; !ok {
		r.uuids = append(r.uuids, provider.UUID)
	}
	r.providers[provider.UUID] = provider

	return provider
}

func (r *MemoryResourceProviderRepository) List() []appplacement.ResourceProvider {
	r.mu.RLock()
	defer r.mu.RUnlock()

	providers := make([]appplacement.ResourceProvider, 0, len(r.uuids))
	for _, uuid := range r.uuids {
		providers = append(providers, r.providers[uuid])
	}

	return providers
}

func (r *MemoryResourceProviderRepository) Get(
	uuid string,
) (appplacement.ResourceProvider, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	provider, ok := r.providers[uuid]
	if !ok {
		return appplacement.ResourceProvider{}, appplacement.ErrResourceProviderNotFound
	}

	return provider, nil
}

func (r *MemoryResourceProviderRepository) Delete(uuid string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.providers[uuid]; !ok {
		return appplacement.ErrResourceProviderNotFound
	}

	delete(r.providers, uuid)
	for index, currentUUID := range r.uuids {
		if currentUUID == uuid {
			r.uuids = append(r.uuids[:index], r.uuids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryResourceProviderRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.uuids = []string{}
	r.providers = map[string]appplacement.ResourceProvider{}
}
