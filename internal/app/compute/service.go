package compute

import "errors"

var ErrFlavorNotFound = errors.New("flavor not found")

type Service struct {
	flavors []Flavor
}

func NewService() *Service {
	return &Service{
		flavors: []Flavor{
			{
				ID:          "1",
				Name:        "m1.small",
				RAM:         2048,
				VCPUs:       1,
				Disk:        20,
				Swap:        0,
				RxTxFactor:  1.0,
				IsPublic:    true,
				Ephemeral:   0,
				Description: "Small test flavor",
				ExtraSpecs:  map[string]string{},
			},
		},
	}
}

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
