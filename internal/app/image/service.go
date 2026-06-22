package image

import (
	"errors"
	"sync"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var ErrImageNotFound = errors.New("image not found")

type Service struct {
	mu     sync.RWMutex
	ids    []string
	images map[string]Image
	clock  clock.Clock
}

func NewService() *Service {
	return NewServiceWithClock(clock.Wall())
}

func NewServiceWithClock(clock clock.Clock) *Service {
	return &Service{
		ids:    []string{},
		images: map[string]Image{},
		clock:  clock,
	}
}

func (s *Service) Create(input CreateImage) Image {
	now := s.clock.Now().UTC()
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

func (s *Service) Get(id string) (Image, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	image, ok := s.images[id]
	if !ok {
		return Image{}, ErrImageNotFound
	}

	return image, nil
}

func (s *Service) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.images[id]; !ok {
		return ErrImageNotFound
	}

	delete(s.images, id)
	for index, currentID := range s.ids {
		if currentID == id {
			s.ids = append(s.ids[:index], s.ids[index+1:]...)
			break
		}
	}

	return nil
}
