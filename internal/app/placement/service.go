package placement

import (
	"errors"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var (
	ErrResourceProviderNotFound = errors.New("resource provider not found")
	ErrInventoryNotFound        = errors.New("inventory not found")
	ErrResourceClassNotFound    = errors.New("resource class not found")
	ErrTraitNotFound            = errors.New("trait not found")
)

type Service struct {
	resourceProviderRepository ResourceProviderRepository
	inventoryRepository        InventoryRepository
	traitRepository            TraitRepository
	aggregateRepository        AggregateRepository
	usageRepository            UsageRepository
	allocationRepository       AllocationRepository
	resourceClassRepository    ResourceClassRepository
	traitCatalogRepository     TraitCatalogRepository
	idGen                      idgen.Generator
}

type Repositories struct {
	ResourceProviders ResourceProviderRepository
	Inventories       InventoryRepository
	ProviderTraits    TraitRepository
	Aggregates        AggregateRepository
	Usages            UsageRepository
	Allocations       AllocationRepository
	ResourceClasses   ResourceClassRepository
	TraitCatalog      TraitCatalogRepository
}

func NewServiceWithRepositories(
	repositories Repositories,
	idGen idgen.Generator,
) *Service {
	return &Service{
		resourceProviderRepository: repositories.ResourceProviders,
		inventoryRepository:        repositories.Inventories,
		traitRepository:            repositories.ProviderTraits,
		aggregateRepository:        repositories.Aggregates,
		usageRepository:            repositories.Usages,
		allocationRepository:       repositories.Allocations,
		resourceClassRepository:    repositories.ResourceClasses,
		traitCatalogRepository:     repositories.TraitCatalog,
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
	s.traitCatalogRepository.Reset()
}
