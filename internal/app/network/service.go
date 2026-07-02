package network

import (
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	mu                      sync.RWMutex
	networkRepository       NetworkRepository
	subnetRepository        SubnetRepository
	portRepository          PortRepository
	securityGroupRepository SecurityGroupRepository
	securityRuleRepository  SecurityGroupRuleRepository
	routerRepository        RouterRepository
	idGen                   idgen.Generator
}

func NewServiceWithRepositories(
	networkRepository NetworkRepository,
	subnetRepository SubnetRepository,
	portRepository PortRepository,
	securityGroupRepository SecurityGroupRepository,
	securityRuleRepository SecurityGroupRuleRepository,
	routerRepository RouterRepository,
	idGen idgen.Generator,
) *Service {
	return &Service{
		networkRepository:       networkRepository,
		subnetRepository:        subnetRepository,
		portRepository:          portRepository,
		securityGroupRepository: securityGroupRepository,
		securityRuleRepository:  securityRuleRepository,
		routerRepository:        routerRepository,
		idGen:                   idGen,
	}
}

func (s *Service) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.networkRepository.Reset()
	s.subnetRepository.Reset()
	s.portRepository.Reset()
	s.securityGroupRepository.Reset()
	s.securityRuleRepository.Reset()
	s.routerRepository.Reset()
}
