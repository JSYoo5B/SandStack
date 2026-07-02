package volume

import (
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	repository         Repository
	snapshotRepository SnapshotRepository
	volumeTypes        []VolumeType
	clock              clock.Clock
	idGen              idgen.Generator
}

func NewServiceWithRuntime(
	repository Repository,
	snapshotRepository SnapshotRepository,
	clock clock.Clock,
	idGen idgen.Generator,
) *Service {
	return &Service{
		repository:         repository,
		snapshotRepository: snapshotRepository,
		volumeTypes: []VolumeType{
			{
				ID:          "default",
				Name:        "__DEFAULT__",
				Description: "Default test volume type",
				ExtraSpecs:  map[string]string{},
				IsPublic:    true,
			},
		},
		clock: clock,
		idGen: idGen,
	}
}

func (s *Service) Reset() {
	s.repository.Reset()
	s.snapshotRepository.Reset()
}
