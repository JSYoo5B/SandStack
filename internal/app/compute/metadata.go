package compute

func (s *Service) ServerMetadata(id string) (map[string]string, error) {
	server, err := s.serverRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return copyMetadata(server.Metadata), nil
}

func (s *Service) ResetServerMetadata(
	id string,
	metadata map[string]string,
) (map[string]string, error) {
	server, err := s.serverRepository.Get(id)
	if err != nil {
		return nil, err
	}

	server.Metadata = copyMetadata(metadata)
	if _, err := s.serverRepository.Update(server); err != nil {
		return nil, err
	}

	return copyMetadata(server.Metadata), nil
}

func (s *Service) UpdateServerMetadata(
	id string,
	metadata map[string]string,
) (map[string]string, error) {
	server, err := s.serverRepository.Get(id)
	if err != nil {
		return nil, err
	}

	if server.Metadata == nil {
		server.Metadata = map[string]string{}
	}
	for key, value := range metadata {
		server.Metadata[key] = value
	}
	if _, err := s.serverRepository.Update(server); err != nil {
		return nil, err
	}

	return copyMetadata(server.Metadata), nil
}

func copyMetadata(metadata map[string]string) map[string]string {
	copied := make(map[string]string, len(metadata))
	for key, value := range metadata {
		copied[key] = value
	}

	return copied
}
