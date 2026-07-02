package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/services"

func (s Service) Services(baseURL string) []services.Service {
	definitions := s.repositories.Services.List()
	result := make([]services.Service, 0, len(definitions))
	for _, entry := range definitions {
		result = append(result, services.Service{
			ID:          entry.ID,
			Name:        entry.Name,
			Description: entry.Description,
			Type:        entry.Type,
			Enabled:     entry.Enabled,
			Links:       map[string]any{},
			Extra:       map[string]any{},
		})
	}

	return result
}

func (s Service) ServiceByID(baseURL, id string) (services.Service, bool) {
	service, err := s.repositories.Services.Get(id)
	if err != nil {
		return services.Service{}, false
	}

	return services.Service{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		Type:        service.Type,
		Enabled:     service.Enabled,
		Links:       map[string]any{},
		Extra:       map[string]any{},
	}, true
}
