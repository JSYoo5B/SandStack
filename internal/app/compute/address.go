package compute

func (s *Service) ServerAddresses(id string) (map[string][]ServerAddress, error) {
	if _, err := s.serverRepository.Get(id); err != nil {
		return nil, err
	}

	return map[string][]ServerAddress{
		"private": {
			{
				Version: 4,
				Address: "10.0.0.10",
			},
		},
	}, nil
}

func (s *Service) ServerAddressesByNetwork(
	id string,
	network string,
) ([]ServerAddress, error) {
	addresses, err := s.ServerAddresses(id)
	if err != nil {
		return nil, err
	}

	return addresses[network], nil
}
