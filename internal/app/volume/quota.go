package volume

import "errors"

var ErrQuotaSetNotFound = errors.New("quota set not found")

func (s *Service) GetQuotaSet(projectID string) QuotaSet {
	quotaSet, err := s.quotaRepository.Get(projectID)
	if err == nil {
		return quotaSet
	}

	return defaultQuotaSet(projectID)
}

func (s *Service) GetDefaultQuotaSet(projectID string) QuotaSet {
	return defaultQuotaSet(projectID)
}

func (s *Service) UpdateQuotaSet(
	projectID string,
	input UpdateQuotaSet,
) QuotaSet {
	quotaSet := s.GetQuotaSet(projectID)
	applyQuotaUpdate(&quotaSet, input)

	return s.quotaRepository.Save(quotaSet)
}

func (s *Service) ResetQuotaSet(projectID string) error {
	return s.quotaRepository.Delete(projectID)
}

func (s *Service) GetQuotaUsageSet(projectID string) QuotaUsageSet {
	quotaSet := s.GetQuotaSet(projectID)
	limits := s.GetLimits()

	return QuotaUsageSet{
		ID: projectID,
		Volumes: quotaUsage(
			limits.TotalVolumesUsed,
			quotaSet.Volumes,
		),
		Snapshots: quotaUsage(
			limits.TotalSnapshotsUsed,
			quotaSet.Snapshots,
		),
		Gigabytes: quotaUsage(
			limits.TotalGigabytesUsed,
			quotaSet.Gigabytes,
		),
		PerVolumeGigabytes: quotaUsage(
			0,
			quotaSet.PerVolumeGigabytes,
		),
		Backups: quotaUsage(
			limits.TotalBackupsUsed,
			quotaSet.Backups,
		),
		BackupGigabytes: quotaUsage(
			limits.TotalBackupGigabytesUsed,
			quotaSet.BackupGigabytes,
		),
		Groups: quotaUsage(0, quotaSet.Groups),
	}
}

func defaultQuotaSet(projectID string) QuotaSet {
	return QuotaSet{
		ID:                 projectID,
		Volumes:            defaultMaxTotalVolumes,
		Snapshots:          defaultMaxTotalSnapshots,
		Gigabytes:          defaultMaxTotalVolumeGigabytes,
		PerVolumeGigabytes: -1,
		Backups:            defaultMaxTotalBackups,
		BackupGigabytes:    defaultMaxTotalBackupGigabytes,
		Groups:             10,
	}
}

func applyQuotaUpdate(quotaSet *QuotaSet, input UpdateQuotaSet) {
	if input.Volumes != nil {
		quotaSet.Volumes = *input.Volumes
	}
	if input.Snapshots != nil {
		quotaSet.Snapshots = *input.Snapshots
	}
	if input.Gigabytes != nil {
		quotaSet.Gigabytes = *input.Gigabytes
	}
	if input.PerVolumeGigabytes != nil {
		quotaSet.PerVolumeGigabytes = *input.PerVolumeGigabytes
	}
	if input.Backups != nil {
		quotaSet.Backups = *input.Backups
	}
	if input.BackupGigabytes != nil {
		quotaSet.BackupGigabytes = *input.BackupGigabytes
	}
	if input.Groups != nil {
		quotaSet.Groups = *input.Groups
	}
}

func quotaUsage(inUse int, limit int) QuotaUsage {
	return QuotaUsage{
		InUse: inUse,
		Limit: limit,
	}
}
