package volume

func (s *Service) ListAvailabilityZones() []AvailabilityZone {
	return []AvailabilityZone{
		{
			Name:      "nova",
			Available: true,
		},
	}
}
