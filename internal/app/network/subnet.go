package network

import (
	"errors"
)

var ErrSubnetNotFound = errors.New("subnet not found")

func (s *Service) ListSubnets() []Subnet {
	s.mu.RLock()
	defer s.mu.RUnlock()

	subnets := make([]Subnet, 0, len(s.subnetIDs))
	for _, id := range s.subnetIDs {
		subnets = append(subnets, s.subnets[id])
	}

	return subnets
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

	s.subnetIDs = append(s.subnetIDs, subnet.ID)
	s.subnets[subnet.ID] = subnet

	network, err := s.networkRepository.Get(subnet.NetworkID)
	if err == nil {
		network.Subnets = append(network.Subnets, subnet.ID)
		_, _ = s.networkRepository.Update(network)
	}

	return subnet
}

func (s *Service) GetSubnet(id string) (Subnet, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	subnet, ok := s.subnets[id]
	if !ok {
		return Subnet{}, ErrSubnetNotFound
	}

	return subnet, nil
}

func (s *Service) DeleteSubnet(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	subnet, ok := s.subnets[id]
	if !ok {
		return ErrSubnetNotFound
	}

	delete(s.subnets, id)
	for index, currentID := range s.subnetIDs {
		if currentID == id {
			s.subnetIDs = append(s.subnetIDs[:index], s.subnetIDs[index+1:]...)
			break
		}
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
