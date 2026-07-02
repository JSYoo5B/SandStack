package placement

import "sync"

type MemoryTraitRepository struct {
	mu     sync.RWMutex
	traits map[string][]string
}

func NewMemoryTraitRepository() *MemoryTraitRepository {
	return &MemoryTraitRepository{
		traits: map[string][]string{},
	}
}

func (r *MemoryTraitRepository) Set(
	resourceProviderUUID string,
	traits []string,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.traits[resourceProviderUUID] = append([]string{}, traits...)
}

func (r *MemoryTraitRepository) Get(resourceProviderUUID string) []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return append([]string{}, r.traits[resourceProviderUUID]...)
}

func (r *MemoryTraitRepository) Delete(resourceProviderUUID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.traits, resourceProviderUUID)
}

func (r *MemoryTraitRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.traits = map[string][]string{}
}
