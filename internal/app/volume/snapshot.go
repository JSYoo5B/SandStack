package volume

import "errors"

var ErrSnapshotNotFound = errors.New("snapshot not found")

func (s *Service) CreateSnapshot(input CreateSnapshot) Snapshot {
	size := 1
	volume, err := s.repository.Get(input.VolumeID)
	if err == nil {
		size = volume.Size
	}

	now := s.clock.Now().UTC().Format(timestampFormat)
	snapshot := Snapshot{
		ID:          "snap-" + s.idGen.Hex(16),
		Name:        input.Name,
		Description: input.Description,
		VolumeID:    input.VolumeID,
		Status:      "available",
		Size:        size,
		Metadata:    input.Metadata,
		Progress:    "100%",
		ProjectID:   "demo",
		UserID:      "demo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if snapshot.Metadata == nil {
		snapshot.Metadata = map[string]string{}
	}

	return s.snapshotRepository.Create(snapshot)
}

func (s *Service) ListSnapshots() []Snapshot {
	return s.snapshotRepository.List()
}

func (s *Service) GetSnapshot(id string) (Snapshot, error) {
	return s.snapshotRepository.Get(id)
}

func (s *Service) DeleteSnapshot(id string) error {
	return s.snapshotRepository.Delete(id)
}
