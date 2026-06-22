package network

import "sync"

type Service struct {
	mu        sync.RWMutex
	ids       []string
	networks  map[string]Network
	subnetIDs []string
	subnets   map[string]Subnet
	portIDs   []string
	ports     map[string]Port
}

func NewService() *Service {
	return &Service{
		ids:       []string{},
		networks:  map[string]Network{},
		subnetIDs: []string{},
		subnets:   map[string]Subnet{},
		portIDs:   []string{},
		ports:     map[string]Port{},
	}
}
