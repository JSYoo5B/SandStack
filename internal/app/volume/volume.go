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

	s.mu.Lock()
	defer s.mu.Unlock()

	s.ids = append(s.ids, volume.ID)
	s.volumes[volume.ID] = volume

	return volume
}

func (s *Service) List() []Volume {
	s.mu.RLock()
	defer s.mu.RUnlock()

	volumes := make([]Volume, 0, len(s.ids))
	for _, id := range s.ids {
		volumes = append(volumes, s.volumes[id])
	}

	return volumes
}

func (s *Service) Get(id string) (Volume, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	volume, ok := s.volumes[id]
	if !ok {
		return Volume{}, ErrVolumeNotFound
	}

	return volume, nil
}

func (s *Service) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.volumes[id]; !ok {
		return ErrVolumeNotFound
	}

	delete(s.volumes, id)
	for index, currentID := range s.ids {
		if currentID == id {
			s.ids = append(s.ids[:index], s.ids[index+1:]...)
			break
		}
	}

	return nil
}
