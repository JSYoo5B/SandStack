package image

import (
	"errors"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var ErrImageNotFound = errors.New("image not found")

type Service struct {
	repository Repository
	clock      clock.Clock
	idGen      idgen.Generator
}

func NewServiceWithRuntime(
	repository Repository,
	clock clock.Clock,
	idGen idgen.Generator,
) *Service {
	return &Service{
		repository: repository,
		clock:      clock,
		idGen:      idGen,
	}
}

func (s *Service) Create(input CreateImage) Image {
	now := s.clock.Now().UTC()
	image := Image{
		ID:              "img-" + s.idGen.Hex(16),
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

	return s.repository.Create(image)
}

func (s *Service) List() []Image {
	return s.repository.List()
}

func (s *Service) Get(id string) (Image, error) {
	return s.repository.Get(id)
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) Reset() {
	s.repository.Reset()
}
