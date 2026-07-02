package network

import (
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	mu                sync.RWMutex
	networkRepository NetworkRepository
	subnetIDs         []string
	subnets           map[string]Subnet
	portIDs           []string
	ports             map[string]Port
	idGen             idgen.Generator
}

func NewService() *Service {
	return NewServiceWithIDGenerator(idgen.Random())
}

func NewServiceWithIDGenerator(idGen idgen.Generator) *Service {
	return NewServiceWithRepositories(
		NewMemoryNetworkRepository(),
		idGen,
	)
}

func NewServiceWithRepositories(
	networkRepository NetworkRepository,
	idGen idgen.Generator,
) *Service {
	return &Service{
		networkRepository: networkRepository,
		subnetIDs:         []string{},
		subnets:           map[string]Subnet{},
		portIDs:           []string{},
		ports:             map[string]Port{},
		idGen:             idGen,
	}
}

func (s *Service) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.networkRepository.Reset()
	s.subnetIDs = []string{}
	s.subnets = map[string]Subnet{}
	s.portIDs = []string{}
	s.ports = map[string]Port{}
}
