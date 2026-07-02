package volume

import appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"

type quotaSetResponse struct {
	QuotaSet quotaSetDocument `json:"quota_set"`
}

type updateQuotaSetRequest struct {
	QuotaSet updateQuotaSetDocument `json:"quota_set"`
}

type quotaSetDocument struct {
	ID                 string `json:"id"`
	Volumes            int    `json:"volumes"`
	Snapshots          int    `json:"snapshots"`
	Gigabytes          int    `json:"gigabytes"`
	PerVolumeGigabytes int    `json:"per_volume_gigabytes"`
	Backups            int    `json:"backups"`
	BackupGigabytes    int    `json:"backup_gigabytes"`
	Groups             int    `json:"groups"`
}

type updateQuotaSetDocument struct {
	Volumes            *int `json:"volumes"`
	Snapshots          *int `json:"snapshots"`
	Gigabytes          *int `json:"gigabytes"`
	PerVolumeGigabytes *int `json:"per_volume_gigabytes"`
	Backups            *int `json:"backups"`
	BackupGigabytes    *int `json:"backup_gigabytes"`
	Groups             *int `json:"groups"`
}

type quotaUsageSetResponse struct {
	QuotaSet quotaUsageSetDocument `json:"quota_set"`
}

type quotaUsageSetDocument struct {
	ID                 string             `json:"id"`
	Volumes            quotaUsageDocument `json:"volumes"`
	Snapshots          quotaUsageDocument `json:"snapshots"`
	Gigabytes          quotaUsageDocument `json:"gigabytes"`
	PerVolumeGigabytes quotaUsageDocument `json:"per_volume_gigabytes"`
	Backups            quotaUsageDocument `json:"backups"`
	BackupGigabytes    quotaUsageDocument `json:"backup_gigabytes"`
	Groups             quotaUsageDocument `json:"groups"`
}

type quotaUsageDocument struct {
	InUse     int `json:"in_use"`
	Reserved  int `json:"reserved"`
	Limit     int `json:"limit"`
	Allocated int `json:"allocated"`
}

func (request updateQuotaSetRequest) updateQuotaSet() appvolume.UpdateQuotaSet {
	return appvolume.UpdateQuotaSet{
		Volumes:            request.QuotaSet.Volumes,
		Snapshots:          request.QuotaSet.Snapshots,
		Gigabytes:          request.QuotaSet.Gigabytes,
		PerVolumeGigabytes: request.QuotaSet.PerVolumeGigabytes,
		Backups:            request.QuotaSet.Backups,
		BackupGigabytes:    request.QuotaSet.BackupGigabytes,
		Groups:             request.QuotaSet.Groups,
	}
}

func toQuotaSetDocument(quotaSet appvolume.QuotaSet) quotaSetDocument {
	return quotaSetDocument{
		ID:                 quotaSet.ID,
		Volumes:            quotaSet.Volumes,
		Snapshots:          quotaSet.Snapshots,
		Gigabytes:          quotaSet.Gigabytes,
		PerVolumeGigabytes: quotaSet.PerVolumeGigabytes,
		Backups:            quotaSet.Backups,
		BackupGigabytes:    quotaSet.BackupGigabytes,
		Groups:             quotaSet.Groups,
	}
}

func toQuotaUsageSetDocument(
	usageSet appvolume.QuotaUsageSet,
) quotaUsageSetDocument {
	return quotaUsageSetDocument{
		ID:                 usageSet.ID,
		Volumes:            toQuotaUsageDocument(usageSet.Volumes),
		Snapshots:          toQuotaUsageDocument(usageSet.Snapshots),
		Gigabytes:          toQuotaUsageDocument(usageSet.Gigabytes),
		PerVolumeGigabytes: toQuotaUsageDocument(usageSet.PerVolumeGigabytes),
		Backups:            toQuotaUsageDocument(usageSet.Backups),
		BackupGigabytes:    toQuotaUsageDocument(usageSet.BackupGigabytes),
		Groups:             toQuotaUsageDocument(usageSet.Groups),
	}
}

func toQuotaUsageDocument(usage appvolume.QuotaUsage) quotaUsageDocument {
	return quotaUsageDocument{
		InUse:     usage.InUse,
		Reserved:  usage.Reserved,
		Limit:     usage.Limit,
		Allocated: usage.Allocated,
	}
}
