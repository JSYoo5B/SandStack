package placement

import (
	"errors"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var (
	ErrResourceProviderNotFound = errors.New("resource provider not found")
	ErrInventoryNotFound        = errors.New("inventory not found")
)

type Service struct {
	resourceProviderRepository ResourceProviderRepository
	inventoryRepository        InventoryRepository
	traitRepository            TraitRepository
	aggregateRepository        AggregateRepository
	usageRepository            UsageRepository
	idGen                      idgen.Generator
}

func NewServiceWithRepositories(
	resourceProviderRepository ResourceProviderRepository,
	inventoryRepository InventoryRepository,
	traitRepository TraitRepository,
	aggregateRepository AggregateRepository,
	usageRepository UsageRepository,
	idGen idgen.Generator,
) *Service {
	return &Service{
		resourceProviderRepository: resourceProviderRepository,
		inventoryRepository:        inventoryRepository,
		traitRepository:            traitRepository,
		aggregateRepository:        aggregateRepository,
		usageRepository:            usageRepository,
		idGen:                      idGen,
	}
}

func (s *Service) CreateResourceProvider(
	input CreateResourceProvider,
) ResourceProvider {
	uuid := input.UUID
	if uuid == "" {
		uuid = "rp-" + s.idGen.Hex(16)
	}

	rootProviderUUID := uuid
	if input.ParentProviderUUID != "" {
		rootProviderUUID = input.ParentProviderUUID
	}

	provider := ResourceProvider{
		UUID:               uuid,
		Name:               input.Name,
		Generation:         0,
		ParentProviderUUID: input.ParentProviderUUID,
		RootProviderUUID:   rootProviderUUID,
	}

	return s.resourceProviderRepository.Create(provider)
}

func (s *Service) ListResourceProviders() []ResourceProvider {
	return s.resourceProviderRepository.List()
}

func (s *Service) GetResourceProvider(uuid string) (ResourceProvider, error) {
	return s.resourceProviderRepository.Get(uuid)
}

func (s *Service) DeleteResourceProvider(uuid string) error {
	err := s.resourceProviderRepository.Delete(uuid)
	if err != nil {
		return err
	}

	s.inventoryRepository.DeleteAll(uuid)
	s.traitRepository.Delete(uuid)
	s.aggregateRepository.Delete(uuid)
	s.usageRepository.Delete(uuid)
	return nil
}

func (s *Service) GetInventories(
	resourceProviderUUID string,
) (Inventories, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return Inventories{}, err
	}

	return Inventories{
		ResourceProviderGeneration: provider.Generation,
		Inventories:                s.inventoryRepository.GetAll(resourceProviderUUID),
	}, nil
}

func (s *Service) UpdateInventories(
	resourceProviderUUID string,
	input UpdateInventories,
) (Inventories, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return Inventories{}, err
	}

	s.inventoryRepository.SetAll(resourceProviderUUID, input.Inventories)

	return Inventories{
		ResourceProviderGeneration: provider.Generation,
		Inventories:                s.inventoryRepository.GetAll(resourceProviderUUID),
	}, nil
}

func (s *Service) GetInventory(
	resourceProviderUUID string,
	resourceClass string,
) (UpdateInventory, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return UpdateInventory{}, err
	}

	inventory, err := s.inventoryRepository.Get(
		resourceProviderUUID,
		resourceClass,
	)
	if err != nil {
		return UpdateInventory{}, err
	}

	return UpdateInventory{
		ResourceProviderGeneration: provider.Generation,
		Inventory:                  inventory,
	}, nil
}

func (s *Service) UpdateInventory(
	resourceProviderUUID string,
	resourceClass string,
	input UpdateInventory,
) (UpdateInventory, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return UpdateInventory{}, err
	}

	inventory := input.Inventory
	inventory.ResourceClass = resourceClass
	s.inventoryRepository.Set(resourceProviderUUID, inventory)

	return UpdateInventory{
		ResourceProviderGeneration: provider.Generation,
		Inventory:                  inventory,
	}, nil
}

func (s *Service) DeleteInventories(resourceProviderUUID string) error {
	_, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return err
	}

	s.inventoryRepository.DeleteAll(resourceProviderUUID)
	return nil
}

func (s *Service) DeleteInventory(
	resourceProviderUUID string,
	resourceClass string,
) error {
	_, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return err
	}

	return s.inventoryRepository.Delete(resourceProviderUUID, resourceClass)
}

func (s *Service) GetTraits(resourceProviderUUID string) (Traits, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return Traits{}, err
	}

	return Traits{
		ResourceProviderGeneration: provider.Generation,
		Traits:                     s.traitRepository.Get(resourceProviderUUID),
	}, nil
}

func (s *Service) UpdateTraits(
	resourceProviderUUID string,
	input UpdateTraits,
) (Traits, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return Traits{}, err
	}

	s.traitRepository.Set(resourceProviderUUID, input.Traits)

	return Traits{
		ResourceProviderGeneration: provider.Generation,
		Traits:                     s.traitRepository.Get(resourceProviderUUID),
	}, nil
}

func (s *Service) DeleteTraits(resourceProviderUUID string) error {
	_, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return err
	}

	s.traitRepository.Delete(resourceProviderUUID)
	return nil
}

func (s *Service) GetAggregates(
	resourceProviderUUID string,
) (Aggregates, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return Aggregates{}, err
	}

	return Aggregates{
		ResourceProviderGeneration: intPtr(provider.Generation),
		Aggregates: s.aggregateRepository.Get(
			resourceProviderUUID,
		),
	}, nil
}

func (s *Service) UpdateAggregates(
	resourceProviderUUID string,
	input UpdateAggregates,
) (Aggregates, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return Aggregates{}, err
	}

	s.aggregateRepository.Set(resourceProviderUUID, input.Aggregates)

	return Aggregates{
		ResourceProviderGeneration: intPtr(provider.Generation),
		Aggregates: s.aggregateRepository.Get(
			resourceProviderUUID,
		),
	}, nil
}

func (s *Service) GetUsages(resourceProviderUUID string) (Usages, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return Usages{}, err
	}

	return Usages{
		ResourceProviderGeneration: provider.Generation,
		Usages:                     s.usageRepository.Get(resourceProviderUUID),
	}, nil
}

func (s *Service) Reset() {
	s.resourceProviderRepository.Reset()
	s.inventoryRepository.Reset()
	s.traitRepository.Reset()
	s.aggregateRepository.Reset()
	s.usageRepository.Reset()
}

func intPtr(value int) *int {
	return &value
}
