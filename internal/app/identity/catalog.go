package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"

func (s Service) Catalog(baseURL string) []tokens.CatalogEntry {
	projectID := s.config.ProjectID

	return []tokens.CatalogEntry{
		s.service("identity", "identity", baseURL+"/identity/v3"),
		s.service("compute", "compute", baseURL+"/compute/v2.1/"+projectID),
		s.service("network", "network", baseURL+"/network/v2.0"),
		s.service("image", "image", baseURL+"/image/v2"),
		s.service("volumev3", "volumev3", baseURL+"/volume/v3/"+projectID),
		s.service("placement", "placement", baseURL+"/placement"),
	}
}

func (s Service) service(id, serviceType, url string) tokens.CatalogEntry {
	return tokens.CatalogEntry{
		ID:   id,
		Name: serviceType,
		Type: serviceType,
		Endpoints: []tokens.Endpoint{
			s.endpoint(id+"-public", "public", url),
			s.endpoint(id+"-internal", "internal", url),
			s.endpoint(id+"-admin", "admin", url),
		},
	}
}

func (s Service) endpoint(id, iface, url string) tokens.Endpoint {
	return tokens.Endpoint{
		ID:        id,
		Interface: iface,
		Region:    s.config.Region,
		RegionID:  s.config.Region,
		URL:       url,
	}
}
