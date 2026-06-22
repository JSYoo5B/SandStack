package compute

import (
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/clock"
)

type Service struct {
	flavors []Flavor
	mu      sync.RWMutex
	ids     []string
	servers map[string]Server
	clock   clock.Clock
}

func NewService() *Service {
	return NewServiceWithClock(clock.Wall())
}

func NewServiceWithClock(clock clock.Clock) *Service {
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
	}
}
