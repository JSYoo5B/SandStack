package network

import "errors"

var ErrFloatingIPNotFound = errors.New("floating IP not found")

func (s *Service) CreateFloatingIP(input CreateFloatingIP) FloatingIP {
	status := "DOWN"
	if input.PortID != "" {
		status = "ACTIVE"
	}

	floatingIP := input.FloatingIP
	if floatingIP == "" {
		floatingIP = "203.0.113.10"
	}

	value := FloatingIP{
		ID:                "fip-" + s.idGen.Hex(16),
		Description:       input.Description,
		FloatingNetworkID: input.FloatingNetworkID,
		FloatingIP:        floatingIP,
		PortID:            input.PortID,
		FixedIP:           input.FixedIP,
		TenantID:          input.ProjectID,
		ProjectID:         input.ProjectID,
		Status:            status,
		Tags:              []string{},
	}

	return s.floatingIPRepository.Create(value)
}

func (s *Service) ListFloatingIPs() []FloatingIP {
	return s.floatingIPRepository.List()
}

func (s *Service) GetFloatingIP(id string) (FloatingIP, error) {
	return s.floatingIPRepository.Get(id)
}

func (s *Service) DeleteFloatingIP(id string) error {
	return s.floatingIPRepository.Delete(id)
}
