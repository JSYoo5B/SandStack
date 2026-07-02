package volume

import (
	"errors"
	"sync"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

const timestampFormat = "2006-01-02T15:04:05.999999"

var ErrVolumeNotFound = errors.New("volume not found")
var ErrVolumeTypeNotFound = errors.New("volume type not found")

type Service struct {
	mu          sync.RWMutex
	ids         []string
	volumes     map[string]Volume
	volumeTypes []VolumeType
}

func NewService() *Service {
	return &Service{
		ids:     []string{},
		volumes: map[string]Volume{},
		volumeTypes: []VolumeType{
			{
				ID:          "default",
				Name:        "__DEFAULT__",
				Description: "Default test volume type",
				ExtraSpecs:  map[string]string{},
				IsPublic:    true,
			},
		},
	}
}

func (s *Service) Create(input CreateVolume) Volume {
	now := time.Now().UTC().Format(timestampFormat)
	volume := Volume{
		ID:          "vol-" + idgen.RandomHex(16),
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

func (s *Service) ListVolumeTypes() []VolumeType {
	volumeTypes := make([]VolumeType, 0, len(s.volumeTypes))
	volumeTypes = append(volumeTypes, s.volumeTypes...)

	return volumeTypes
}

func (s *Service) GetVolumeType(id string) (VolumeType, error) {
	for _, volumeType := range s.volumeTypes {
		if volumeType.ID == id {
			return volumeType, nil
		}
	}

	return VolumeType{}, ErrVolumeTypeNotFound
}
