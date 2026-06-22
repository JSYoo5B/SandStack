package image

import (
	"sync"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	mu     sync.RWMutex
	ids    []string
	images map[string]Image
}

func NewService() *Service {
	return &Service{
		ids:    []string{},
		images: map[string]Image{},
	}
}

func (s *Service) Create(input CreateImage) Image {
	now := time.Now().UTC()
	image := Image{
		ID:              "img-" + idgen.RandomHex(16),
		Name:            input.Name,
		Status:          "queued",
		ContainerFormat: input.ContainerFormat,
		DiskFormat:      input.DiskFormat,
		MinDisk:         input.MinDisk,
		MinRAM:          input.MinRAM,
		Protected:       false,
		Visibility:      "private",
		Tags:            input.Tags,
		CreatedAt:       now.Format(time.RFC3339),
		UpdatedAt:       now.Format(time.RFC3339),
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.ids = append(s.ids, image.ID)
	s.images[image.ID] = image

	return image
}

func (s *Service) List() []Image {
	s.mu.RLock()
	defer s.mu.RUnlock()

	images := make([]Image, 0, len(s.ids))
	for _, id := range s.ids {
		images = append(images, s.images[id])
	}

	return images
}
