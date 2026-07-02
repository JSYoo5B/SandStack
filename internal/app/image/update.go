package image

type PatchImage struct {
	Name       *string
	MinDisk    *int
	MinRAM     *int
	Protected  *bool
	Visibility *string
	Tags       []string
	HasTags    bool
}

func (s *Service) Update(id string, patch PatchImage) (Image, error) {
	image, err := s.repository.Get(id)
	if err != nil {
		return Image{}, err
	}

	if patch.Name != nil {
		image.Name = *patch.Name
	}
	if patch.MinDisk != nil {
		image.MinDisk = *patch.MinDisk
	}
	if patch.MinRAM != nil {
		image.MinRAM = *patch.MinRAM
	}
	if patch.Protected != nil {
		image.Protected = *patch.Protected
	}
	if patch.Visibility != nil {
		image.Visibility = *patch.Visibility
	}
	if patch.HasTags {
		image.Tags = patch.Tags
	}

	return s.repository.Update(image)
}
