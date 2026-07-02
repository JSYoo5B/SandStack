package placement

import "strings"

func (s *Service) CreateTrait(name string) (created bool) {
	if _, err := s.traitCatalogRepository.Get(name); err == nil {
		return false
	}

	s.traitCatalogRepository.Create(name)
	return true
}

func (s *Service) ListTraits(nameFilter string) []string {
	traits := s.traitCatalogRepository.List()
	if nameFilter == "" {
		return traits
	}

	if after, ok := strings.CutPrefix(nameFilter, "startswith:"); ok {
		return filterTraits(traits, func(trait string) bool {
			return strings.HasPrefix(trait, after)
		})
	}

	if after, ok := strings.CutPrefix(nameFilter, "in:"); ok {
		allowed := map[string]bool{}
		for _, name := range strings.Split(after, ",") {
			allowed[name] = true
		}
		return filterTraits(traits, func(trait string) bool {
			return allowed[trait]
		})
	}

	return filterTraits(traits, func(trait string) bool {
		return trait == nameFilter
	})
}

func (s *Service) GetTrait(name string) (string, error) {
	return s.traitCatalogRepository.Get(name)
}

func (s *Service) DeleteTrait(name string) error {
	return s.traitCatalogRepository.Delete(name)
}

func filterTraits(traits []string, keep func(string) bool) []string {
	filtered := []string{}
	for _, trait := range traits {
		if keep(trait) {
			filtered = append(filtered, trait)
		}
	}

	return filtered
}
