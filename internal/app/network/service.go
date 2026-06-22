package network

import (
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	mu       sync.RWMutex
	ids      []string
	networks map[string]Network
}

func NewService() *Service {
	return &Service{
		ids:      []string{},
		networks: map[string]Network{},
	}
}

func (s *Service) Create(input CreateNetwork) Network {
	adminStateUp := true
	if input.AdminStateUp != nil {
		adminStateUp = *input.AdminStateUp
	}

	network := Network{
		ID:           "net-" + idgen.RandomHex(16),
		Name:         input.Name,
		Description:  input.Description,
		AdminStateUp: adminStateUp,
		Status:       "ACTIVE",
		Subnets:      []string{},
		TenantID:     input.ProjectID,
		ProjectID:    input.ProjectID,
		Shared:       input.Shared,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.ids = append(s.ids, network.ID)
	s.networks[network.ID] = network

	return network
}

func (s *Service) List() []Network {
	s.mu.RLock()
	defer s.mu.RUnlock()

	networks := make([]Network, 0, len(s.ids))
	for _, id := range s.ids {
		networks = append(networks, s.networks[id])
	}

	return networks
}
