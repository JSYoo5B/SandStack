package identity

import appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"

type versionsResponse struct {
	Versions versionValues `json:"versions"`
}

type versionValues struct {
	Values []versionDocument `json:"values"`
}

type versionResponse struct {
	Version versionDocument `json:"version"`
}

type versionDocument struct {
	ID         string        `json:"id"`
	Status     string        `json:"status"`
	Updated    string        `json:"updated"`
	Links      []versionLink `json:"links"`
	MediaTypes []mediaType   `json:"media-types"`
}

type versionLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type mediaType struct {
	Base string `json:"base"`
	Type string `json:"type"`
}

func toVersionDocument(version appidentity.VersionDocument) versionDocument {
	links := make([]versionLink, 0, len(version.Links))
	for _, link := range version.Links {
		links = append(links, versionLink{
			Rel:  link.Rel,
			Href: link.Href,
		})
	}

	mediaTypes := make([]mediaType, 0, len(version.MediaTypes))
	for _, media := range version.MediaTypes {
		mediaTypes = append(mediaTypes, mediaType{
			Base: media.Base,
			Type: media.Type,
		})
	}

	return versionDocument{
		ID:         version.ID,
		Status:     version.Status,
		Updated:    version.Updated,
		Links:      links,
		MediaTypes: mediaTypes,
	}
}
