package compute

import "errors"

var ErrFlavorNotFound = errors.New("flavor not found")

func (s *Service) ListFlavors() []Flavor {
	flavors := make([]Flavor, 0, len(s.flavors))
	flavors = append(flavors, s.flavors...)

	return flavors
}

func (s *Service) GetFlavor(id string) (Flavor, error) {
	for _, flavor := range s.flavors {
		if flavor.ID == id {
			return flavor, nil
		}
	}

	return Flavor{}, ErrFlavorNotFound
}

func (s *Service) ListFlavorExtraSpecs(id string) (map[string]string, error) {
	flavor, err := s.GetFlavor(id)
	if err != nil {
		return nil, err
	}

	return copyStringMap(flavor.ExtraSpecs), nil
}

func (s *Service) GetFlavorExtraSpec(
	id string,
	key string,
) (map[string]string, error) {
	extraSpecs, err := s.ListFlavorExtraSpecs(id)
	if err != nil {
		return nil, err
	}

	value, ok := extraSpecs[key]
	if !ok {
		return nil, ErrFlavorNotFound
	}

	return map[string]string{key: value}, nil
}

func (s *Service) CreateFlavorExtraSpecs(
	id string,
	extraSpecs map[string]string,
) (map[string]string, error) {
	index, err := s.findFlavorIndex(id)
	if err != nil {
		return nil, err
	}
	if s.flavors[index].ExtraSpecs == nil {
		s.flavors[index].ExtraSpecs = map[string]string{}
	}
	for key, value := range extraSpecs {
		s.flavors[index].ExtraSpecs[key] = value
	}

	return copyStringMap(s.flavors[index].ExtraSpecs), nil
}

func (s *Service) UpdateFlavorExtraSpec(
	id string,
	key string,
	extraSpecs map[string]string,
) (map[string]string, error) {
	value, ok := extraSpecs[key]
	if !ok {
		return nil, ErrFlavorNotFound
	}

	_, err := s.CreateFlavorExtraSpecs(
		id,
		map[string]string{key: value},
	)
	if err != nil {
		return nil, err
	}

	return map[string]string{key: value}, nil
}

func (s *Service) DeleteFlavorExtraSpec(id string, key string) error {
	index, err := s.findFlavorIndex(id)
	if err != nil {
		return err
	}

	delete(s.flavors[index].ExtraSpecs, key)

	return nil
}

func (s *Service) findFlavorIndex(id string) (int, error) {
	for index, flavor := range s.flavors {
		if flavor.ID == id {
			return index, nil
		}
	}

	return 0, ErrFlavorNotFound
}

func copyStringMap(values map[string]string) map[string]string {
	copied := make(map[string]string, len(values))
	for key, value := range values {
		copied[key] = value
	}

	return copied
}
