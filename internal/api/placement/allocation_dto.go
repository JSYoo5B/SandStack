package placement

import appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"

type allocationsDocument struct {
	ResourceProviderGeneration int                           `json:"resource_provider_generation"`
	Allocations                map[string]allocationDocument `json:"allocations"`
}

type allocationDocument struct {
	Resources map[string]int `json:"resources"`
}

func toAllocationsDocument(
	allocations appplacement.Allocations,
) allocationsDocument {
	documents := map[string]allocationDocument{}
	for consumerID, allocation := range allocations.Allocations {
		documents[consumerID] = allocationDocument{
			Resources: allocation.Resources,
		}
	}

	return allocationsDocument{
		ResourceProviderGeneration: allocations.ResourceProviderGeneration,
		Allocations:                documents,
	}
}
