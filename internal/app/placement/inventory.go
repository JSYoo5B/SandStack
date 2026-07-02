package placement

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
