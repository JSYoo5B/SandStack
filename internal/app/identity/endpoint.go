package identity

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/endpoints"
)

func (s Service) Endpoints(baseURL string) []endpoints.Endpoint {
	definitions := s.repositories.Endpoints.List()
	result := make([]endpoints.Endpoint, 0, len(definitions))
	for _, endpoint := range definitions {
		service, err := s.repositories.Services.Get(endpoint.ServiceID)
		if err != nil {
			continue
		}
		result = append(result, toEndpoint(baseURL, service, endpoint))
	}

	return result
}

func (s Service) EndpointByID(
	baseURL string,
	id string,
) (endpoints.Endpoint, bool) {
	endpoint, err := s.repositories.Endpoints.Get(id)
	if err != nil {
		return endpoints.Endpoint{}, false
	}
	service, err := s.repositories.Services.Get(endpoint.ServiceID)
	if err != nil {
		return endpoints.Endpoint{}, false
	}

	return toEndpoint(baseURL, service, endpoint), true
}

func toEndpoint(
	baseURL string,
	service ServiceDefinition,
	endpoint EndpointDefinition,
) endpoints.Endpoint {
	return endpoints.Endpoint{
		ID:           endpoint.ID,
		Availability: gophercloud.Availability(endpoint.Interface),
		Name:         service.Name,
		Region:       endpoint.Region,
		ServiceID:    endpoint.ServiceID,
		URL:          baseURL + endpoint.Path,
		Enabled:      endpoint.Enabled,
		Description:  endpoint.Description,
	}
}
