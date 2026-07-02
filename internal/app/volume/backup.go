package volume

import "errors"

var ErrBackupNotFound = errors.New("backup not found")

func (s *Service) CreateBackup(input CreateBackup) Backup {
	size := 1
	volume, err := s.repository.Get(input.VolumeID)
	if err == nil {
		size = volume.Size
	}

	now := s.clock.Now().UTC().Format(timestampFormat)
	backup := Backup{
		ID:               "backup-" + s.idGen.Hex(16),
		Name:             input.Name,
		Description:      input.Description,
		VolumeID:         input.VolumeID,
		SnapshotID:       input.SnapshotID,
		Status:           "available",
		Size:             size,
		ObjectCount:      1,
		Container:        input.Container,
		IsIncremental:    input.Incremental,
		ProjectID:        "demo",
		Metadata:         input.Metadata,
		AvailabilityZone: input.AvailabilityZone,
		CreatedAt:        now,
		UpdatedAt:        now,
		DataTimestamp:    now,
	}
	if backup.Metadata == nil {
		backup.Metadata = map[string]string{}
	}

	return s.backupRepository.Create(backup)
}

func (s *Service) ListBackups() []Backup {
	return s.backupRepository.List()
}

func (s *Service) GetBackup(id string) (Backup, error) {
	return s.backupRepository.Get(id)
}

func (s *Service) DeleteBackup(id string) error {
	return s.backupRepository.Delete(id)
}
