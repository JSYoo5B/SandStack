package network

import (
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	mu                sync.RWMutex
	networkRepository NetworkRepository
	subnetRepository  SubnetRepository
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
		NewMemorySubnetRepository(),
		idGen,
	)
}

func NewServiceWithRepositories(
	networkRepository NetworkRepository,
	subnetRepository SubnetRepository,
	idGen idgen.Generator,
) *Service {
	return &Service{
		networkRepository: networkRepository,
		subnetRepository:  subnetRepository,
		portIDs:           []string{},
		ports:             map[string]Port{},
		idGen:             idGen,
	}
}

func (s *Service) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.networkRepository.Reset()
	s.subnetRepository.Reset()
	s.portIDs = []string{}
	s.ports = map[string]Port{}
}
