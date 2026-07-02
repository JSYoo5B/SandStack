package placement

import appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"

type traitsDocument struct {
	ResourceProviderGeneration int      `json:"resource_provider_generation"`
	Traits                     []string `json:"traits"`
}

func toTraitsDocument(traits appplacement.Traits) traitsDocument {
	return traitsDocument{
		ResourceProviderGeneration: traits.ResourceProviderGeneration,
		Traits:                     traits.Traits,
	}
}
