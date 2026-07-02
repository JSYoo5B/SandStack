package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"

type catalogService struct {
	id          string
	serviceType string
	path        string
}

func (s Service) Catalog(baseURL string) []tokens.CatalogEntry {
	services := s.repositories.Services.List()
	catalog := make([]tokens.CatalogEntry, 0, len(services))
	for _, service := range services {
		catalog = append(catalog, s.catalogEntry(baseURL, service))
	}

	return catalog
}

func (s Service) catalogEntry(
	baseURL string,
	service ServiceDefinition,
) tokens.CatalogEntry {
	endpoints := s.repositories.Endpoints.ListByServiceID(service.ID)
	tokenEndpoints := make([]tokens.Endpoint, 0, len(endpoints))
	for _, endpoint := range endpoints {
		tokenEndpoints = append(tokenEndpoints, tokens.Endpoint{
			ID:        endpoint.ID,
			Interface: endpoint.Interface,
			Region:    endpoint.Region,
			RegionID:  endpoint.Region,
			URL:       baseURL + endpoint.Path,
		})
	}

	return tokens.CatalogEntry{
		ID:        service.ID,
		Name:      service.Name,
		Type:      service.Type,
		Endpoints: tokenEndpoints,
	}
}

func (s Service) defaultServices() []ServiceDefinition {
	return []ServiceDefinition{
		s.defaultService("identity", "identity"),
		s.defaultService("compute", "compute"),
		s.defaultService("network", "network"),
		s.defaultService("image", "image"),
		s.defaultService("volumev3", "volumev3"),
		s.defaultService("placement", "placement"),
	}
}

func (s Service) defaultService(id string, serviceType string) ServiceDefinition {
	return ServiceDefinition{
		ID:          id,
		Name:        serviceType,
		Type:        serviceType,
		Description: "SandStack " + serviceType + " service",
		Enabled:     true,
	}
}

func (s Service) defaultEndpoints() []EndpointDefinition {
	projectID := s.config.ProjectID
	services := []catalogService{
		{id: "identity", serviceType: "identity", path: "/identity/v3"},
		{id: "compute", serviceType: "compute", path: "/compute/v2.1/" + projectID},
		{id: "network", serviceType: "network", path: "/network/v2.0"},
		{id: "image", serviceType: "image", path: "/image/v2"},
		{id: "volumev3", serviceType: "volumev3", path: "/volume/v3/" + projectID},
		{id: "placement", serviceType: "placement", path: "/placement"},
	}

	endpoints := []EndpointDefinition{}
	for _, service := range services {
		endpoints = append(
			endpoints,
			s.defaultEndpoint(service.id, "public", service.path),
			s.defaultEndpoint(service.id, "internal", service.path),
			s.defaultEndpoint(service.id, "admin", service.path),
		)
	}

	return endpoints
}

func (s Service) defaultEndpoint(
	serviceID string,
	iface string,
	path string,
) EndpointDefinition {
	return EndpointDefinition{
		ID:          serviceID + "-" + iface,
		ServiceID:   serviceID,
		Interface:   iface,
		Region:      s.config.Region,
		Path:        path,
		Enabled:     true,
		Description: "SandStack " + serviceID + " endpoint",
	}
}
