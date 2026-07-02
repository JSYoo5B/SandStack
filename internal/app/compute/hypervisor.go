package compute

import "errors"

var ErrHypervisorNotFound = errors.New("hypervisor not found")

func (s *Service) ListHypervisors() []Hypervisor {
	return []Hypervisor{s.defaultHypervisor()}
}

func (s *Service) GetHypervisor(id string) (Hypervisor, error) {
	hypervisor := s.defaultHypervisor()
	if hypervisor.ID != id {
		return Hypervisor{}, ErrHypervisorNotFound
	}

	return hypervisor, nil
}

func (s *Service) defaultHypervisor() Hypervisor {
	return Hypervisor{
		ID:                 "1",
		Status:             "enabled",
		State:              "up",
		DiskAvailableLeast: 80,
		HostIP:             "127.0.0.1",
		FreeDiskGB:         80,
		FreeRAMMB:          512000,
		Hostname:           "sandstack",
		Type:               "sandstack",
		Version:            1,
		LocalGB:            100,
		LocalGBUsed:        20,
		MemoryMB:           512000,
		MemoryMBUsed:       0,
		ServiceID:          "compute-service-1",
		VCPUs:              200,
	}
}
