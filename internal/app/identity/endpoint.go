package identity

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/endpoints"
)

func (s Service) Endpoints(baseURL string) []endpoints.Endpoint {
	catalog := s.Catalog(baseURL)
	result := []endpoints.Endpoint{}
	for _, entry := range catalog {
		for _, endpoint := range entry.Endpoints {
			result = append(result, endpoints.Endpoint{
				ID:           endpoint.ID,
				Availability: gophercloud.Availability(endpoint.Interface),
				Name:         entry.Name,
				Region:       endpoint.Region,
				ServiceID:    entry.ID,
				URL:          endpoint.URL,
				Enabled:      true,
				Description:  "SandStack " + entry.Name + " endpoint",
			})
		}
	}

	return result
}

func (s Service) EndpointByID(
	baseURL string,
	id string,
) (endpoints.Endpoint, bool) {
	for _, endpoint := range s.Endpoints(baseURL) {
		if endpoint.ID == id {
			return endpoint, true
		}
	}

	return endpoints.Endpoint{}, false
}
