package placement

import "sync"

type MemoryAggregateRepository struct {
	mu         sync.RWMutex
	aggregates map[string][]string
}

func NewMemoryAggregateRepository() *MemoryAggregateRepository {
	return &MemoryAggregateRepository{
		aggregates: map[string][]string{},
	}
}

func (r *MemoryAggregateRepository) Set(
	resourceProviderUUID string,
	aggregates []string,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.aggregates[resourceProviderUUID] = append([]string{}, aggregates...)
}

func (r *MemoryAggregateRepository) Get(
	resourceProviderUUID string,
) []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return append([]string{}, r.aggregates[resourceProviderUUID]...)
}

func (r *MemoryAggregateRepository) Delete(resourceProviderUUID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.aggregates, resourceProviderUUID)
}

func (r *MemoryAggregateRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.aggregates = map[string][]string{}
}
