package compute

import (
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	flavors []Flavor
	mu      sync.RWMutex
	ids     []string
	servers map[string]Server
	clock   clock.Clock
	idGen   idgen.Generator
}

func NewService() *Service {
	return NewServiceWithRuntime(clock.Wall(), idgen.Random())
}

func NewServiceWithClock(clock clock.Clock) *Service {
	return NewServiceWithRuntime(clock, idgen.Random())
}

func NewServiceWithRuntime(
	clock clock.Clock,
	idGen idgen.Generator,
) *Service {
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
		clock:   clock,
		idGen:   idGen,
	}
}

func (s *Service) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ids = []string{}
	s.servers = map[string]Server{}
}
