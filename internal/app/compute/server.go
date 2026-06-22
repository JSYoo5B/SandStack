package compute

import (
	"errors"
	"time"
)

var ErrServerNotFound = errors.New("server not found")

const serverTimestampFormat = time.RFC3339

func (s *Service) ListServers() []Server {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.activateServersLocked()

	servers := make([]Server, 0, len(s.ids))
	for _, id := range s.ids {
		servers = append(servers, s.servers[id])
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

	s.mu.Lock()
	defer s.mu.Unlock()

	s.ids = append(s.ids, server.ID)
	s.servers[server.ID] = server

	return server
}

func (s *Service) GetServer(id string) (Server, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	server, ok := s.servers[id]
	if !ok {
		return Server{}, ErrServerNotFound
	}

	server = s.activateServerLocked(server)
	s.servers[id] = server

	return server, nil
}

func (s *Service) DeleteServer(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.servers[id]; !ok {
		return ErrServerNotFound
	}

	delete(s.servers, id)
	for index, currentID := range s.ids {
		if currentID == id {
			s.ids = append(s.ids[:index], s.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (s *Service) activateServersLocked() {
	for _, id := range s.ids {
		s.servers[id] = s.activateServerLocked(s.servers[id])
	}
}

func (s *Service) activateServerLocked(server Server) Server {
	if server.Status != "BUILD" {
		return server
	}

	server.Status = "ACTIVE"
	server.Progress = 100
	server.UpdatedAt = s.clock.Now().UTC().Format(serverTimestampFormat)

	return server
}
