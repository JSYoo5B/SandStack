package placement

import (
	"sync"

	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
)

type MemoryInventoryRepository struct {
	mu          sync.RWMutex
	inventories map[string]map[string]appplacement.Inventory
}

func NewMemoryInventoryRepository() *MemoryInventoryRepository {
	return &MemoryInventoryRepository{
		inventories: map[string]map[string]appplacement.Inventory{},
	}
}

func (r *MemoryInventoryRepository) SetAll(
	resourceProviderUUID string,
	inventories map[string]appplacement.Inventory,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.inventories[resourceProviderUUID] = copyInventories(inventories)
}

func (r *MemoryInventoryRepository) GetAll(
	resourceProviderUUID string,
) map[string]appplacement.Inventory {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return copyInventories(r.inventories[resourceProviderUUID])
}

func (r *MemoryInventoryRepository) Set(
	resourceProviderUUID string,
	inventory appplacement.Inventory,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.inventories[resourceProviderUUID]; !ok {
		r.inventories[resourceProviderUUID] = map[string]appplacement.Inventory{}
	}
	r.inventories[resourceProviderUUID][inventory.ResourceClass] = inventory
}

func (r *MemoryInventoryRepository) Get(
	resourceProviderUUID string,
	resourceClass string,
) (appplacement.Inventory, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	inventories, ok := r.inventories[resourceProviderUUID]
	if !ok {
		return appplacement.Inventory{}, appplacement.ErrInventoryNotFound
	}

	inventory, ok := inventories[resourceClass]
	if !ok {
		return appplacement.Inventory{}, appplacement.ErrInventoryNotFound
	}

	return inventory, nil
}

func (r *MemoryInventoryRepository) DeleteAll(resourceProviderUUID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.inventories, resourceProviderUUID)
}

func (r *MemoryInventoryRepository) Delete(
	resourceProviderUUID string,
	resourceClass string,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	inventories, ok := r.inventories[resourceProviderUUID]
	if !ok {
		return appplacement.ErrInventoryNotFound
	}
	if _, ok := inventories[resourceClass]; !ok {
		return appplacement.ErrInventoryNotFound
	}

	delete(inventories, resourceClass)
	return nil
}

func (r *MemoryInventoryRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.inventories = map[string]map[string]appplacement.Inventory{}
}

func copyInventories(
	inventories map[string]appplacement.Inventory,
) map[string]appplacement.Inventory {
	copied := map[string]appplacement.Inventory{}
	for resourceClass, inventory := range inventories {
		inventory.ResourceClass = resourceClass
		copied[resourceClass] = inventory
	}

	return copied
}
