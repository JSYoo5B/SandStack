package volume

type UpdateVolume struct {
	Name        *string
	Description *string
	Metadata    map[string]string
}

func (s *Service) Update(id string, input UpdateVolume) (Volume, error) {
	volume, err := s.repository.Get(id)
	if err != nil {
		return Volume{}, err
	}

	if input.Name != nil {
		volume.Name = *input.Name
	}
	if input.Description != nil {
		volume.Description = *input.Description
	}
	if input.Metadata != nil {
		volume.Metadata = input.Metadata
	}

	return s.repository.Update(volume)
}
