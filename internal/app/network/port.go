package network

import (
	"errors"
)

var ErrPortNotFound = errors.New("port not found")

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

	id := "port-" + s.idGen.Hex(16)
	port := Port{
		ID:           id,
		NetworkID:    input.NetworkID,
		Name:         input.Name,
		Description:  input.Description,
		AdminStateUp: adminStateUp,
		Status:       "DOWN",
		MACAddress:   "fa:16:3e:" + s.idGen.Hex(6),
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

func (s *Service) GetPort(id string) (Port, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	port, ok := s.ports[id]
	if !ok {
		return Port{}, ErrPortNotFound
	}

	return port, nil
}

func (s *Service) DeletePort(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.ports[id]; !ok {
		return ErrPortNotFound
	}

	delete(s.ports, id)
	for index, currentID := range s.portIDs {
		if currentID == id {
			s.portIDs = append(s.portIDs[:index], s.portIDs[index+1:]...)
			break
		}
	}

	return nil
}
