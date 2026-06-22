package volume

import (
	"sync"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

const timestampFormat = "2006-01-02T15:04:05.999999"

type Service struct {
	mu      sync.RWMutex
	ids     []string
	volumes map[string]Volume
}

func NewService() *Service {
	return &Service{
		ids:     []string{},
		volumes: map[string]Volume{},
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
