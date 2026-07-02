package placement

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
