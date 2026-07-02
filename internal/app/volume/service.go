package volume

import (
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	repository           Repository
	snapshotRepository   SnapshotRepository
	transferRepository   TransferRepository
	backupRepository     BackupRepository
	attachmentRepository AttachmentRepository
	volumeTypes          []VolumeType
	clock                clock.Clock
	idGen                idgen.Generator
}

func NewServiceWithRuntime(
	repository Repository,
	snapshotRepository SnapshotRepository,
	transferRepository TransferRepository,
	backupRepository BackupRepository,
	attachmentRepository AttachmentRepository,
	clock clock.Clock,
	idGen idgen.Generator,
) *Service {
	return &Service{
		repository:           repository,
		snapshotRepository:   snapshotRepository,
		transferRepository:   transferRepository,
		backupRepository:     backupRepository,
		attachmentRepository: attachmentRepository,
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
	s.transferRepository.Reset()
	s.backupRepository.Reset()
	s.attachmentRepository.Reset()
}
