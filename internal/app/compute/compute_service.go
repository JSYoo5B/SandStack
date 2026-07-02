package compute

func (s *Service) ListComputeServices() []ComputeService {
	return []ComputeService{
		{
			ID:        "compute-service-1",
			Binary:    "nova-compute",
			Host:      "sandstack",
			State:     "up",
			Status:    "enabled",
			UpdatedAt: "2026-07-03T00:00:00.000000",
			Zone:      "nova",
		},
		{
			ID:        "scheduler-service-1",
			Binary:    "nova-scheduler",
			Host:      "sandstack",
			State:     "up",
			Status:    "enabled",
			UpdatedAt: "2026-07-03T00:00:00.000000",
			Zone:      "internal",
		},
	}
}
