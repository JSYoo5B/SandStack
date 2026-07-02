package placement

import appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"

type aggregatesDocument struct {
	ResourceProviderGeneration *int     `json:"resource_provider_generation,omitempty"`
	Aggregates                 []string `json:"aggregates"`
}

func toAggregatesDocument(
	aggregates appplacement.Aggregates,
) aggregatesDocument {
	return aggregatesDocument{
		ResourceProviderGeneration: aggregates.ResourceProviderGeneration,
		Aggregates:                 aggregates.Aggregates,
	}
}
