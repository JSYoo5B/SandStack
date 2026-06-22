package network

import (
	"errors"
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var ErrNetworkNotFound = errors.New("network not found")
var ErrSubnetNotFound = errors.New("subnet not found")

type Service struct {
	mu        sync.RWMutex
	ids       []string
	networks  map[string]Network
	subnetIDs []string
	subnets   map[string]Subnet
	portIDs   []string
	ports     map[string]Port
}

func NewService() *Service {
	return &Service{
		ids:       []string{},
		networks:  map[string]Network{},
		subnetIDs: []string{},
		subnets:   map[string]Subnet{},
		portIDs:   []string{},
		ports:     map[string]Port{},
	}
}

func (s *Service) Create(input CreateNetwork) Network {
	adminStateUp := true
	if input.AdminStateUp != nil {
		adminStateUp = *input.AdminStateUp
	}

	network := Network{
		ID:           "net-" + idgen.RandomHex(16),
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
		ID:             "subnet-" + idgen.RandomHex(16),
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

	network, ok := s.networks[subnet.NetworkID]
	if ok {
		network.Subnets = append(network.Subnets, subnet.ID)
		s.networks[subnet.NetworkID] = network
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

	network, ok := s.networks[subnet.NetworkID]
	if ok {
		network.Subnets = removeString(network.Subnets, id)
		s.networks[subnet.NetworkID] = network
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

func (s *Service) ListPorts() []Port {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ports := make([]Port, 0, len(s.portIDs))
	for _, id := range s.portIDs {
		ports = append(ports, s.ports[id])
	}

	return ports
}

func (s *Service) CreatePort(input CreatePort) Port {
	adminStateUp := true
	if input.AdminStateUp != nil {
		adminStateUp = *input.AdminStateUp
	}

	id := "port-" + idgen.RandomHex(16)
	port := Port{
		ID:           id,
		NetworkID:    input.NetworkID,
		Name:         input.Name,
		Description:  input.Description,
		AdminStateUp: adminStateUp,
		Status:       "DOWN",
		MACAddress:   "fa:16:3e:" + idgen.RandomHex(6),
		FixedIPs:     input.FixedIPs,
		TenantID:     input.ProjectID,
		ProjectID:    input.ProjectID,
		DeviceID:     input.DeviceID,
		DeviceOwner:  input.DeviceOwner,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.portIDs = append(s.portIDs, port.ID)
	s.ports[port.ID] = port

	return port
}
