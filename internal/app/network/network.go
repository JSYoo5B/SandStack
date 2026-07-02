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

	return s.networkRepository.Create(network)
}

func (s *Service) List() []Network {
	return s.networkRepository.List()
}

func (s *Service) Get(id string) (Network, error) {
	return s.networkRepository.Get(id)
}

func (s *Service) Delete(id string) error {
	return s.networkRepository.Delete(id)
}
