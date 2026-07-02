package placement

import appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"

type allocationCandidatesDocument struct {
	AllocationRequests []allocationRequestDocument        `json:"allocation_requests"`
	ProviderSummaries  map[string]providerSummaryDocument `json:"provider_summaries"`
}

type allocationRequestDocument struct {
	Allocations map[string]allocationDocument `json:"allocations"`
}

type providerSummaryDocument struct {
	Resources          map[string]providerSummaryResourceDocument `json:"resources"`
	Traits             []string                                   `json:"traits,omitempty"`
	ParentProviderUUID string                                     `json:"parent_provider_uuid,omitempty"`
	RootProviderUUID   string                                     `json:"root_provider_uuid,omitempty"`
}

type providerSummaryResourceDocument struct {
	Capacity int `json:"capacity"`
	Used     int `json:"used"`
}

func toAllocationCandidatesDocument(
	candidates appplacement.AllocationCandidates,
) allocationCandidatesDocument {
	return allocationCandidatesDocument{
		AllocationRequests: toAllocationRequestDocuments(
			candidates.AllocationRequests,
		),
		ProviderSummaries: toProviderSummaryDocuments(
			candidates.ProviderSummaries,
		),
	}
}

func toAllocationRequestDocuments(
	requests []appplacement.AllocationRequest,
) []allocationRequestDocument {
	documents := make([]allocationRequestDocument, 0, len(requests))
	for _, request := range requests {
		allocations := map[string]allocationDocument{}
		for providerUUID, allocation := range request.Allocations {
			allocations[providerUUID] = allocationDocument{
				Resources: allocation.Resources,
			}
		}
		documents = append(documents, allocationRequestDocument{
			Allocations: allocations,
		})
	}

	return documents
}

func toProviderSummaryDocuments(
	summaries map[string]appplacement.ProviderSummary,
) map[string]providerSummaryDocument {
	documents := map[string]providerSummaryDocument{}
	for providerUUID, summary := range summaries {
		documents[providerUUID] = providerSummaryDocument{
			Resources: toProviderSummaryResourceDocuments(
				summary.Resources,
			),
			Traits:             summary.Traits,
			ParentProviderUUID: summary.ParentProviderUUID,
			RootProviderUUID:   summary.RootProviderUUID,
		}
	}

	return documents
}

func toProviderSummaryResourceDocuments(
	resources map[string]appplacement.ProviderSummaryResource,
) map[string]providerSummaryResourceDocument {
	documents := map[string]providerSummaryResourceDocument{}
	for resourceClass, resource := range resources {
		documents[resourceClass] = providerSummaryResourceDocument{
			Capacity: resource.Capacity,
			Used:     resource.Used,
		}
	}

	return documents
}
