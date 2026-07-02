package volume

const (
	defaultMaxTotalVolumes         = 1000
	defaultMaxTotalSnapshots       = 1000
	defaultMaxTotalVolumeGigabytes = 100000
	defaultMaxTotalBackups         = 1000
	defaultMaxTotalBackupGigabytes = 100000
)

func (s *Service) GetLimits() Limits {
	volumes := s.repository.List()
	snapshots := s.snapshotRepository.List()
	backups := s.backupRepository.List()

	return Limits{
		MaxTotalVolumes:          defaultMaxTotalVolumes,
		MaxTotalSnapshots:        defaultMaxTotalSnapshots,
		MaxTotalVolumeGigabytes:  defaultMaxTotalVolumeGigabytes,
		MaxTotalBackups:          defaultMaxTotalBackups,
		MaxTotalBackupGigabytes:  defaultMaxTotalBackupGigabytes,
		TotalVolumesUsed:         len(volumes),
		TotalGigabytesUsed:       totalVolumeSize(volumes),
		TotalSnapshotsUsed:       len(snapshots),
		TotalBackupsUsed:         len(backups),
		TotalBackupGigabytesUsed: totalBackupSize(backups),
	}
}

func totalVolumeSize(volumes []Volume) int {
	total := 0
	for _, volume := range volumes {
		total += volume.Size
	}

	return total
}

func totalBackupSize(backups []Backup) int {
	total := 0
	for _, backup := range backups {
		total += backup.Size
	}

	return total
}
