package placement

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
