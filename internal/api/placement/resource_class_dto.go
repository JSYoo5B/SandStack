package placement

import appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"

type createResourceClassRequest struct {
	Name string `json:"name"`
}

type resourceClassListResponse struct {
	ResourceClasses []resourceClassDocument `json:"resource_classes"`
}

type resourceClassDocument struct {
	Name  string                      `json:"name"`
	Links []resourceClassLinkDocument `json:"links"`
}

type resourceClassLinkDocument struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

func toResourceClassDocuments(
	classes []appplacement.ResourceClass,
) []resourceClassDocument {
	documents := make([]resourceClassDocument, 0, len(classes))
	for _, class := range classes {
		documents = append(documents, toResourceClassDocument(class))
	}

	return documents
}

func toResourceClassDocument(
	class appplacement.ResourceClass,
) resourceClassDocument {
	return resourceClassDocument{
		Name: class.Name,
		Links: []resourceClassLinkDocument{
			{
				Href: "/placement/resource_classes/" + class.Name,
				Rel:  "self",
			},
		},
	}
}
