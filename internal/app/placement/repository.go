package placement

type ResourceProviderRepository interface {
	Create(provider ResourceProvider) ResourceProvider
	List() []ResourceProvider
	Get(uuid string) (ResourceProvider, error)
	Delete(uuid string) error
	Reset()
}

type InventoryRepository interface {
	SetAll(resourceProviderUUID string, inventories map[string]Inventory)
	GetAll(resourceProviderUUID string) map[string]Inventory
	Set(resourceProviderUUID string, inventory Inventory)
	Get(resourceProviderUUID string, resourceClass string) (Inventory, error)
	DeleteAll(resourceProviderUUID string)
	Delete(resourceProviderUUID string, resourceClass string) error
	Reset()
}

type TraitRepository interface {
	Set(resourceProviderUUID string, traits []string)
	Get(resourceProviderUUID string) []string
	Delete(resourceProviderUUID string)
	Reset()
}

type AggregateRepository interface {
	Set(resourceProviderUUID string, aggregates []string)
	Get(resourceProviderUUID string) []string
	Delete(resourceProviderUUID string)
	Reset()
}

type UsageRepository interface {
	Set(resourceProviderUUID string, usages map[string]int)
	Get(resourceProviderUUID string) map[string]int
	Delete(resourceProviderUUID string)
	Reset()
}

type AllocationRepository interface {
	Set(resourceProviderUUID string, allocations map[string]Allocation)
	Get(resourceProviderUUID string) map[string]Allocation
	Delete(resourceProviderUUID string)
	Reset()
}

type ResourceClassRepository interface {
	Create(resourceClass ResourceClass) ResourceClass
	List() []ResourceClass
	Get(name string) (ResourceClass, error)
	Delete(name string) error
	Reset()
}

type TraitCatalogRepository interface {
	Create(name string)
	List() []string
	Get(name string) (string, error)
	Delete(name string) error
	Reset()
}
