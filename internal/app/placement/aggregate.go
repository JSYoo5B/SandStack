package placement

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

func intPtr(value int) *int {
	return &value
}
