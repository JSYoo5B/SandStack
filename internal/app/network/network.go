package network

import (
	"errors"
)

var ErrNetworkNotFound = errors.New("network not found")

func (s *Service) Create(input CreateNetwork) Network {
	adminStateUp := true
	if input.AdminStateUp != nil {
		adminStateUp = *input.AdminStateUp
	}

	network := Network{
		ID:           "net-" + s.idGen.Hex(16),
		Name:         input.Name,
		Description:  input.Description,
		AdminStateUp: adminStateUp,
		Status:       "ACTIVE",
		Subnets:      []string{},
		TenantID:     input.ProjectID,
		ProjectID:    input.ProjectID,
		Shared:       input.Shared,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.ids = append(s.ids, network.ID)
	s.networks[network.ID] = network

	return network
}

func (s *Service) List() []Network {
	s.mu.RLock()
	defer s.mu.RUnlock()

	networks := make([]Network, 0, len(s.ids))
	for _, id := range s.ids {
		networks = append(networks, s.networks[id])
	}

	return networks
}

func (s *Service) Get(id string) (Network, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	network, ok := s.networks[id]
	if !ok {
		return Network{}, ErrNetworkNotFound
	}

	return network, nil
}

func (s *Service) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.networks[id]; !ok {
		return ErrNetworkNotFound
	}

	delete(s.networks, id)
	for index, currentID := range s.ids {
		if currentID == id {
			s.ids = append(s.ids[:index], s.ids[index+1:]...)
			break
		}
	}

	return nil
}
