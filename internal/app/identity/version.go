package identity

func (s Service) Version(baseURL string) VersionDocument {
	return VersionDocument{
		ID:      "v3.14",
		Status:  "stable",
		Updated: "2020-04-07T00:00:00Z",
		Links: []VersionLink{
			{
				Rel:  "self",
				Href: baseURL + "/identity/v3/",
			},
		},
		MediaTypes: []MediaType{
			{
				Base: "application/json",
				Type: "application/vnd.openstack.identity-v3+json",
			},
		},
	}
}
