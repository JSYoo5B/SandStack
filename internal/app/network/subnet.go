package network

import (
	"errors"
)

var ErrSubnetNotFound = errors.New("subnet not found")

func (s *Service) ListSubnets() []Subnet {
	return s.subnetRepository.List()
}

func (s *Service) CreateSubnet(input CreateSubnet) Subnet {
	enableDHCP := true
	if input.EnableDHCP != nil {
		enableDHCP = *input.EnableDHCP
	}

	subnet := Subnet{
		ID:             "subnet-" + s.idGen.Hex(16),
		NetworkID:      input.NetworkID,
		Name:           input.Name,
		Description:    input.Description,
		IPVersion:      input.IPVersion,
		CIDR:           input.CIDR,
		GatewayIP:      input.GatewayIP,
		DNSNameservers: input.DNSNameservers,
		EnableDHCP:     enableDHCP,
		TenantID:       input.ProjectID,
		ProjectID:      input.ProjectID,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.subnetRepository.Create(subnet)

	network, err := s.networkRepository.Get(subnet.NetworkID)
	if err == nil {
		network.Subnets = append(network.Subnets, subnet.ID)
		_, _ = s.networkRepository.Update(network)
	}

	return subnet
}

func (s *Service) GetSubnet(id string) (Subnet, error) {
	return s.subnetRepository.Get(id)
}

func (s *Service) DeleteSubnet(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	subnet, err := s.subnetRepository.Get(id)
	if err != nil {
		return err
	}

	if err := s.subnetRepository.Delete(id); err != nil {
		return err
	}

	network, err := s.networkRepository.Get(subnet.NetworkID)
	if err == nil {
		network.Subnets = removeString(network.Subnets, id)
		_, _ = s.networkRepository.Update(network)
	}

	return nil
}

func removeString(values []string, target string) []string {
	for index, value := range values {
		if value == target {
			return append(values[:index], values[index+1:]...)
		}
	}

	return values
}
