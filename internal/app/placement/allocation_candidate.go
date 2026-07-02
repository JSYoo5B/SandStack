package placement

func (s *Service) ListAllocationCandidates(
	query AllocationCandidateQuery,
) AllocationCandidates {
	candidates := AllocationCandidates{
		AllocationRequests: []AllocationRequest{},
		ProviderSummaries:  map[string]ProviderSummary{},
	}

	if len(query.Resources) == 0 {
		return candidates
	}

	for _, provider := range s.resourceProviderRepository.List() {
		inventories := s.inventoryRepository.GetAll(provider.UUID)
		usages := s.usageRepository.Get(provider.UUID)

		if !canSatisfyResources(inventories, usages, query.Resources) {
			continue
		}

		candidates.AllocationRequests = append(
			candidates.AllocationRequests,
			AllocationRequest{
				Allocations: map[string]Allocation{
					provider.UUID: {
						Resources: copyIntMap(query.Resources),
					},
				},
			},
		)
		candidates.ProviderSummaries[provider.UUID] = ProviderSummary{
			Resources:          toProviderSummaryResources(inventories, usages),
			Traits:             s.traitRepository.Get(provider.UUID),
			ParentProviderUUID: provider.ParentProviderUUID,
			RootProviderUUID:   provider.RootProviderUUID,
		}
	}

	return candidates
}

func canSatisfyResources(
	inventories map[string]Inventory,
	usages map[string]int,
	requested map[string]int,
) bool {
	for resourceClass, amount := range requested {
		inventory, ok := inventories[resourceClass]
		if !ok {
			return false
		}
		if capacity(inventory, usages[resourceClass]) < amount {
			return false
		}
	}

	return true
}

func toProviderSummaryResources(
	inventories map[string]Inventory,
	usages map[string]int,
) map[string]ProviderSummaryResource {
	resources := map[string]ProviderSummaryResource{}
	for resourceClass, inventory := range inventories {
		used := usages[resourceClass]
		resources[resourceClass] = ProviderSummaryResource{
			Capacity: capacity(inventory, used),
			Used:     used,
		}
	}

	return resources
}

func capacity(inventory Inventory, used int) int {
	available := inventory.Total - inventory.Reserved
	return int(float32(available)*inventory.AllocationRatio) - used
}

func copyIntMap(values map[string]int) map[string]int {
	copied := map[string]int{}
	for key, value := range values {
		copied[key] = value
	}

	return copied
}
