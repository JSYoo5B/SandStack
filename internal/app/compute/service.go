package compute

import (
	"errors"
	"sync"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var ErrFlavorNotFound = errors.New("flavor not found")

type Service struct {
	flavors []Flavor
	mu      sync.RWMutex
	ids     []string
	servers map[string]Server
}

func NewService() *Service {
	return &Service{
		flavors: []Flavor{
			{
				ID:          "1",
				Name:        "m1.small",
				RAM:         2048,
				VCPUs:       1,
				Disk:        20,
				Swap:        0,
				RxTxFactor:  1.0,
				IsPublic:    true,
				Ephemeral:   0,
				Description: "Small test flavor",
				ExtraSpecs:  map[string]string{},
			},
		},
		ids:     []string{},
		servers: map[string]Server{},
	}
}

func (s *Service) ListFlavors() []Flavor {
	flavors := make([]Flavor, 0, len(s.flavors))
	flavors = append(flavors, s.flavors...)

	return flavors
}

func (s *Service) GetFlavor(id string) (Flavor, error) {
	for _, flavor := range s.flavors {
		if flavor.ID == id {
			return flavor, nil
		}
	}

	return Flavor{}, ErrFlavorNotFound
}

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
