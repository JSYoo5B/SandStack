package placement

func (s *Service) CreateResourceClass(name string) ResourceClass {
	return s.resourceClassRepository.Create(ResourceClass{Name: name})
}

func (s *Service) EnsureResourceClass(name string) ResourceClass {
	if resourceClass, err := s.resourceClassRepository.Get(name); err == nil {
		return resourceClass
	}

	return s.CreateResourceClass(name)
}

func (s *Service) ListResourceClasses() []ResourceClass {
	return s.resourceClassRepository.List()
}

func (s *Service) GetResourceClass(name string) (ResourceClass, error) {
	return s.resourceClassRepository.Get(name)
}

func (s *Service) DeleteResourceClass(name string) error {
	return s.resourceClassRepository.Delete(name)
}
