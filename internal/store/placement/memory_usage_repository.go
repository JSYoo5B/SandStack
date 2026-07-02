package placement

import "sync"

type MemoryUsageRepository struct {
	mu     sync.RWMutex
	usages map[string]map[string]int
}

func NewMemoryUsageRepository() *MemoryUsageRepository {
	return &MemoryUsageRepository{
		usages: map[string]map[string]int{},
	}
}

func (r *MemoryUsageRepository) Set(
	resourceProviderUUID string,
	usages map[string]int,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.usages[resourceProviderUUID] = copyUsages(usages)
}

func (r *MemoryUsageRepository) Get(resourceProviderUUID string) map[string]int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return copyUsages(r.usages[resourceProviderUUID])
}

func (r *MemoryUsageRepository) Delete(resourceProviderUUID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.usages, resourceProviderUUID)
}

func (r *MemoryUsageRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.usages = map[string]map[string]int{}
}

func copyUsages(usages map[string]int) map[string]int {
	copied := map[string]int{}
	for resourceClass, used := range usages {
		copied[resourceClass] = used
	}

	return copied
}
