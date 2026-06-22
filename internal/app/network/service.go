package network

import (
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	mu        sync.RWMutex
	ids       []string
	networks  map[string]Network
	subnetIDs []string
	subnets   map[string]Subnet
	portIDs   []string
	ports     map[string]Port
	idGen     idgen.Generator
}

func NewService() *Service {
	return NewServiceWithIDGenerator(idgen.Random())
}

func NewServiceWithIDGenerator(idGen idgen.Generator) *Service {
	return &Service{
		ids:       []string{},
		networks:  map[string]Network{},
		subnetIDs: []string{},
		subnets:   map[string]Subnet{},
		portIDs:   []string{},
		ports:     map[string]Port{},
		idGen:     idGen,
	}
}
