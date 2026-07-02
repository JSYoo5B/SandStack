package compute

import (
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	flavors           []Flavor
	serverRepository  ServerRepository
	keyPairRepository KeyPairRepository
	clock             clock.Clock
	idGen             idgen.Generator
}

func NewServiceWithRuntime(
	serverRepository ServerRepository,
	keyPairRepository KeyPairRepository,
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
		serverRepository:  serverRepository,
		keyPairRepository: keyPairRepository,
		clock:             clock,
		idGen:             idGen,
	}
}

func (s *Service) Reset() {
	s.serverRepository.Reset()
	s.keyPairRepository.Reset()
}
