package placement

import appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"

type createResourceProviderRequest struct {
	Name               string `json:"name"`
	UUID               string `json:"uuid"`
	ParentProviderUUID string `json:"parent_provider_uuid"`
}

type resourceProviderListResponse struct {
	ResourceProviders []resourceProviderDocument `json:"resource_providers"`
}

type resourceProviderDocument struct {
	Generation         int                            `json:"generation"`
	UUID               string                         `json:"uuid"`
	Links              []resourceProviderLinkDocument `json:"links"`
	Name               string                         `json:"name"`
	ParentProviderUUID string                         `json:"parent_provider_uuid"`
	RootProviderUUID   string                         `json:"root_provider_uuid"`
}

type resourceProviderLinkDocument struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

func toResourceProviderDocuments(
	providers []appplacement.ResourceProvider,
) []resourceProviderDocument {
	documents := make([]resourceProviderDocument, 0, len(providers))
	for _, provider := range providers {
		documents = append(documents, toResourceProviderDocument(provider))
	}

	return documents
}

func toResourceProviderDocument(
	provider appplacement.ResourceProvider,
) resourceProviderDocument {
	return resourceProviderDocument{
		Generation: provider.Generation,
		UUID:       provider.UUID,
		Links: []resourceProviderLinkDocument{
			{
				Href: "/placement/resource_providers/" + provider.UUID,
				Rel:  "self",
			},
		},
		Name:               provider.Name,
		ParentProviderUUID: provider.ParentProviderUUID,
		RootProviderUUID:   provider.RootProviderUUID,
	}
}
