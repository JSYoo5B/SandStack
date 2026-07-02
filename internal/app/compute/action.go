package compute

func (s *Service) StartServer(id string) error {
	server, err := s.serverRepository.Get(id)
	if err != nil {
		return err
	}

	server.Status = "ACTIVE"
	server.Progress = 100
	_, err = s.serverRepository.Update(server)
	return err
}

func (s *Service) StopServer(id string) error {
	server, err := s.serverRepository.Get(id)
	if err != nil {
		return err
	}

	server.Status = "SHUTOFF"
	server.Progress = 0
	_, err = s.serverRepository.Update(server)
	return err
}
