package placement

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
