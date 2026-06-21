package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"

type catalogService struct {
	id          string
	serviceType string
	path        string
}

func (s Service) Catalog(baseURL string) []tokens.CatalogEntry {
	projectID := s.config.ProjectID

	catalog := []tokens.CatalogEntry{
		s.service("identity", "identity", baseURL+"/identity/v3"),
	}

	optionalServices := []catalogService{
		{id: "compute", serviceType: "compute", path: "/compute/v2.1/" + projectID},
		{id: "network", serviceType: "network", path: "/network/v2.0"},
		{id: "image", serviceType: "image", path: "/image/v2"},
		{id: "volumev3", serviceType: "volumev3", path: "/volume/v3/" + projectID},
		{id: "placement", serviceType: "placement", path: "/placement"},
	}

	for _, service := range optionalServices {
		catalog = append(catalog, s.service(
			service.id,
			service.serviceType,
			baseURL+service.path,
		))
	}

	return catalog
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
