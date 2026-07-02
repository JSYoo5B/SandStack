package placement

import (
	"errors"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var (
	ErrResourceProviderNotFound = errors.New("resource provider not found")
	ErrInventoryNotFound        = errors.New("inventory not found")
	ErrResourceClassNotFound    = errors.New("resource class not found")
)

type Service struct {
	resourceProviderRepository ResourceProviderRepository
	inventoryRepository        InventoryRepository
	traitRepository            TraitRepository
	aggregateRepository        AggregateRepository
	usageRepository            UsageRepository
	allocationRepository       AllocationRepository
	resourceClassRepository    ResourceClassRepository
	idGen                      idgen.Generator
}

func NewServiceWithRepositories(
	resourceProviderRepository ResourceProviderRepository,
	inventoryRepository InventoryRepository,
	traitRepository TraitRepository,
	aggregateRepository AggregateRepository,
	usageRepository UsageRepository,
	allocationRepository AllocationRepository,
	resourceClassRepository ResourceClassRepository,
	idGen idgen.Generator,
) *Service {
	return &Service{
		resourceProviderRepository: resourceProviderRepository,
		inventoryRepository:        inventoryRepository,
		traitRepository:            traitRepository,
		aggregateRepository:        aggregateRepository,
		usageRepository:            usageRepository,
		allocationRepository:       allocationRepository,
		resourceClassRepository:    resourceClassRepository,
		idGen:                      idGen,
	}
}

func (s *Service) Reset() {
	s.resourceProviderRepository.Reset()
	s.inventoryRepository.Reset()
	s.traitRepository.Reset()
	s.aggregateRepository.Reset()
	s.usageRepository.Reset()
	s.allocationRepository.Reset()
	s.resourceClassRepository.Reset()
}
