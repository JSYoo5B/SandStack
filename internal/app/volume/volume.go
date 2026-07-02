package volume

import (
	"errors"
)

const timestampFormat = "2006-01-02T15:04:05.999999"

var ErrVolumeNotFound = errors.New("volume not found")

func (s *Service) Create(input CreateVolume) Volume {
	now := s.clock.Now().UTC().Format(timestampFormat)
	volume := Volume{
		ID:          "vol-" + s.idGen.Hex(16),
		Status:      "creating",
		Size:        input.Size,
		Name:        input.Name,
		Description: input.Description,
		VolumeType:  input.VolumeType,
		Metadata:    input.Metadata,
		CreatedAt:   now,
		UpdatedAt:   now,
		Bootable:    "false",
	}

	return s.repository.Create(volume)
}

func (s *Service) List() []Volume {
	volumes := s.repository.List()
	for index, volume := range volumes {
		volumes[index] = s.makeVolumeAvailable(volume)
	}

	return volumes
}

func (s *Service) Get(id string) (Volume, error) {
	volume, err := s.repository.Get(id)
	if err != nil {
		return Volume{}, err
	}

	return s.makeVolumeAvailable(volume), nil
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) makeVolumeAvailable(volume Volume) Volume {
	if volume.Status != "creating" {
		return volume
	}

	volume.Status = "available"
	volume.UpdatedAt = s.clock.Now().UTC().Format(timestampFormat)

	_, _ = s.repository.Update(volume)

	return volume
}
