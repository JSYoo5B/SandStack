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
	idGen                      idgen.Generator
}

func NewServiceWithRepositories(
	resourceProviderRepository ResourceProviderRepository,
	inventoryRepository InventoryRepository,
	idGen idgen.Generator,
) *Service {
	return &Service{
		resourceProviderRepository: resourceProviderRepository,
		inventoryRepository:        inventoryRepository,
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

func (s *Service) Reset() {
	s.resourceProviderRepository.Reset()
	s.inventoryRepository.Reset()
}
