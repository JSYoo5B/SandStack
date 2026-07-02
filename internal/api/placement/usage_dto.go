package placement

import appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"

type usagesDocument struct {
	ResourceProviderGeneration int            `json:"resource_provider_generation"`
	Usages                     map[string]int `json:"usages"`
}

func toUsagesDocument(usages appplacement.Usages) usagesDocument {
	return usagesDocument{
		ResourceProviderGeneration: usages.ResourceProviderGeneration,
		Usages:                     usages.Usages,
	}
}
