package compute

import "errors"

var ErrMetadatumNotFound = errors.New("server metadatum not found")

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

func (s *Service) ServerMetadatum(
	id string,
	key string,
) (map[string]string, error) {
	metadata, err := s.ServerMetadata(id)
	if err != nil {
		return nil, err
	}

	value, ok := metadata[key]
	if !ok {
		return nil, ErrMetadatumNotFound
	}

	return map[string]string{key: value}, nil
}

func (s *Service) SetServerMetadatum(
	id string,
	key string,
	value string,
) (map[string]string, error) {
	return s.UpdateServerMetadata(id, map[string]string{key: value})
}

func (s *Service) DeleteServerMetadatum(id string, key string) error {
	server, err := s.serverRepository.Get(id)
	if err != nil {
		return err
	}

	delete(server.Metadata, key)
	_, err = s.serverRepository.Update(server)
	return err
}

func copyMetadata(metadata map[string]string) map[string]string {
	copied := make(map[string]string, len(metadata))
	for key, value := range metadata {
		copied[key] = value
	}

	return copied
}
