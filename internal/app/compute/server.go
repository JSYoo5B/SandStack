package compute

import (
	"errors"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var ErrServerNotFound = errors.New("server not found")

func (s *Service) ListServers() []Server {
	s.mu.RLock()
	defer s.mu.RUnlock()

	servers := make([]Server, 0, len(s.ids))
	for _, id := range s.ids {
		servers = append(servers, s.servers[id])
	}

	return servers
}

func (s *Service) CreateServer(input CreateServer) Server {
	now := time.Now().UTC().Format(time.RFC3339)
	server := Server{
		ID:        "srv-" + idgen.RandomHex(16),
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
	s.mu.RLock()
	defer s.mu.RUnlock()

	server, ok := s.servers[id]
	if !ok {
		return Server{}, ErrServerNotFound
	}

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
