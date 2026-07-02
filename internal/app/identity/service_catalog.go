package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/services"

func (s Service) Services(baseURL string) []services.Service {
	catalog := s.Catalog(baseURL)
	result := make([]services.Service, 0, len(catalog))
	for _, entry := range catalog {
		result = append(result, services.Service{
			ID:          entry.ID,
			Name:        entry.Name,
			Description: "SandStack " + entry.Name + " service",
			Type:        entry.Type,
			Enabled:     true,
			Links:       map[string]any{},
			Extra:       map[string]any{},
		})
	}

	return result
}

func (s Service) ServiceByID(baseURL, id string) (services.Service, bool) {
	for _, service := range s.Services(baseURL) {
		if service.ID == id {
			return service, true
		}
	}

	return services.Service{}, false
}
