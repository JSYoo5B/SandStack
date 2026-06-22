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
