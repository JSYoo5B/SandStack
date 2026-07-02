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
