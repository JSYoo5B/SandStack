package image

import (
	"errors"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var ErrImageNotFound = errors.New("image not found")
var ErrTaskNotFound = errors.New("task not found")

type Service struct {
	repository       Repository
	dataRepository   DataRepository
	memberRepository MemberRepository
	taskRepository   TaskRepository
	clock            clock.Clock
	idGen            idgen.Generator
}

func NewServiceWithRuntime(
	repository Repository,
	dataRepository DataRepository,
	memberRepository MemberRepository,
	taskRepository TaskRepository,
	clock clock.Clock,
	idGen idgen.Generator,
) *Service {
	return &Service{
		repository:       repository,
		dataRepository:   dataRepository,
		memberRepository: memberRepository,
		taskRepository:   taskRepository,
		clock:            clock,
		idGen:            idGen,
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
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	s.dataRepository.Delete(id)
	return nil
}

func (s *Service) Reset() {
	s.repository.Reset()
	s.dataRepository.Reset()
	s.memberRepository.Reset()
	s.taskRepository.Reset()
}
