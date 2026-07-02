package network

import (
	"errors"
)

var ErrPortNotFound = errors.New("port not found")

func (s *Service) ListPorts() []Port {
	return s.portRepository.List()
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

	return s.portRepository.Create(port)
}

func (s *Service) GetPort(id string) (Port, error) {
	return s.portRepository.Get(id)
}

func (s *Service) DeletePort(id string) error {
	return s.portRepository.Delete(id)
}
