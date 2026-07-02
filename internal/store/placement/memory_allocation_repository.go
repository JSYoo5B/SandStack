package placement

import (
	"sync"

	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
)

type MemoryAllocationRepository struct {
	mu          sync.RWMutex
	allocations map[string]map[string]appplacement.Allocation
}

func NewMemoryAllocationRepository() *MemoryAllocationRepository {
	return &MemoryAllocationRepository{
		allocations: map[string]map[string]appplacement.Allocation{},
	}
}

func (r *MemoryAllocationRepository) Set(
	resourceProviderUUID string,
	allocations map[string]appplacement.Allocation,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.allocations[resourceProviderUUID] = copyAllocations(allocations)
}

func (r *MemoryAllocationRepository) Get(
	resourceProviderUUID string,
) map[string]appplacement.Allocation {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return copyAllocations(r.allocations[resourceProviderUUID])
}

func (r *MemoryAllocationRepository) Delete(resourceProviderUUID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.allocations, resourceProviderUUID)
}

func (r *MemoryAllocationRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.allocations = map[string]map[string]appplacement.Allocation{}
}

func copyAllocations(
	allocations map[string]appplacement.Allocation,
) map[string]appplacement.Allocation {
	copied := map[string]appplacement.Allocation{}
	for consumerID, allocation := range allocations {
		copied[consumerID] = appplacement.Allocation{
			Resources: copyResources(allocation.Resources),
		}
	}

	return copied
}

func copyResources(resources map[string]int) map[string]int {
	copied := map[string]int{}
	for resourceClass, amount := range resources {
		copied[resourceClass] = amount
	}

	return copied
}
