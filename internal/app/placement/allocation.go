package placement

func (s *Service) GetAllocations(
	resourceProviderUUID string,
) (Allocations, error) {
	provider, err := s.resourceProviderRepository.Get(resourceProviderUUID)
	if err != nil {
		return Allocations{}, err
	}

	return Allocations{
		ResourceProviderGeneration: provider.Generation,
		Allocations: s.allocationRepository.Get(
			resourceProviderUUID,
		),
	}, nil
}
