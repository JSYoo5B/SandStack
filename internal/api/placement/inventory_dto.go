package placement

import appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"

type inventoriesDocument struct {
	ResourceProviderGeneration int                          `json:"resource_provider_generation"`
	Inventories                map[string]inventoryDocument `json:"inventories"`
}

type inventoryDocument struct {
	ResourceProviderGeneration int     `json:"resource_provider_generation,omitempty"`
	AllocationRatio            float32 `json:"allocation_ratio"`
	MaxUnit                    int     `json:"max_unit"`
	MinUnit                    int     `json:"min_unit"`
	Reserved                   int     `json:"reserved"`
	StepSize                   int     `json:"step_size"`
	Total                      int     `json:"total"`
}

func toInventoriesDocument(
	inventories appplacement.Inventories,
) inventoriesDocument {
	documents := map[string]inventoryDocument{}
	for resourceClass, inventory := range inventories.Inventories {
		documents[resourceClass] = toInventoryDocument(inventory)
	}

	return inventoriesDocument{
		ResourceProviderGeneration: inventories.ResourceProviderGeneration,
		Inventories:                documents,
	}
}

func toInventoryDocument(
	inventory appplacement.Inventory,
) inventoryDocument {
	return inventoryDocument{
		AllocationRatio: inventory.AllocationRatio,
		MaxUnit:         inventory.MaxUnit,
		MinUnit:         inventory.MinUnit,
		Reserved:        inventory.Reserved,
		StepSize:        inventory.StepSize,
		Total:           inventory.Total,
	}
}

func toInventoryWithGenerationDocument(
	inventory appplacement.UpdateInventory,
) inventoryDocument {
	document := toInventoryDocument(inventory.Inventory)
	document.ResourceProviderGeneration = inventory.ResourceProviderGeneration

	return document
}

func toAppInventories(
	document inventoriesDocument,
) map[string]appplacement.Inventory {
	inventories := map[string]appplacement.Inventory{}
	for resourceClass, inventory := range document.Inventories {
		inventories[resourceClass] = toAppInventory(resourceClass, inventory)
	}

	return inventories
}

func toAppInventory(
	resourceClass string,
	document inventoryDocument,
) appplacement.Inventory {
	return appplacement.Inventory{
		ResourceClass:   resourceClass,
		AllocationRatio: document.AllocationRatio,
		MaxUnit:         document.MaxUnit,
		MinUnit:         document.MinUnit,
		Reserved:        document.Reserved,
		StepSize:        document.StepSize,
		Total:           document.Total,
	}
}
