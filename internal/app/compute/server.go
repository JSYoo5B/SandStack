package compute

import (
	"errors"
	"time"
)

var ErrServerNotFound = errors.New("server not found")

const serverTimestampFormat = time.RFC3339

func (s *Service) ListServers() []Server {
	servers := s.serverRepository.List()
	for index, server := range servers {
		servers[index] = s.activateServer(server)
	}

	return servers
}

func (s *Service) CreateServer(input CreateServer) Server {
	now := s.clock.Now().UTC().Format(serverTimestampFormat)
	server := Server{
		ID:        "srv-" + s.idGen.Hex(16),
		Name:      input.Name,
		ImageID:   input.ImageID,
		FlavorID:  input.FlavorID,
		TenantID:  "demo",
		UserID:    "admin",
		Status:    "BUILD",
		Progress:  0,
		CreatedAt: now,
		UpdatedAt: now,
		Metadata:  input.Metadata,
	}

	return s.serverRepository.Create(server)
}

func (s *Service) GetServer(id string) (Server, error) {
	server, err := s.serverRepository.Get(id)
	if err != nil {
		return Server{}, err
	}

	return s.activateServer(server), nil
}

func (s *Service) DeleteServer(id string) error {
	return s.serverRepository.Delete(id)
}

func (s *Service) activateServer(server Server) Server {
	if server.Status != "BUILD" {
		return server
	}

	server.Status = "ACTIVE"
	server.Progress = 100
	server.UpdatedAt = s.clock.Now().UTC().Format(serverTimestampFormat)

	_, _ = s.serverRepository.Update(server)

	return server
}
