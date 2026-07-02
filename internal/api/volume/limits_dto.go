package volume

import appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"

type limitsResponse struct {
	Limits limitsDocument `json:"limits"`
}

type limitsDocument struct {
	Absolute absoluteLimitsDocument `json:"absolute"`
	Rate     []rateLimitDocument    `json:"rate"`
}

type absoluteLimitsDocument struct {
	MaxTotalVolumes          int `json:"maxTotalVolumes"`
	MaxTotalSnapshots        int `json:"maxTotalSnapshots"`
	MaxTotalVolumeGigabytes  int `json:"maxTotalVolumeGigabytes"`
	MaxTotalBackups          int `json:"maxTotalBackups"`
	MaxTotalBackupGigabytes  int `json:"maxTotalBackupGigabytes"`
	TotalVolumesUsed         int `json:"totalVolumesUsed"`
	TotalGigabytesUsed       int `json:"totalGigabytesUsed"`
	TotalSnapshotsUsed       int `json:"totalSnapshotsUsed"`
	TotalBackupsUsed         int `json:"totalBackupsUsed"`
	TotalBackupGigabytesUsed int `json:"totalBackupGigabytesUsed"`
}

type rateLimitDocument struct {
	Regex string                  `json:"regex"`
	URI   string                  `json:"uri"`
	Limit []rateLimitItemDocument `json:"limit"`
}

type rateLimitItemDocument struct {
	Verb          string `json:"verb"`
	NextAvailable string `json:"next-available"`
	Unit          string `json:"unit"`
	Value         int    `json:"value"`
	Remaining     int    `json:"remaining"`
}

func toLimitsDocument(limits appvolume.Limits) limitsDocument {
	return limitsDocument{
		Absolute: absoluteLimitsDocument{
			MaxTotalVolumes:          limits.MaxTotalVolumes,
			MaxTotalSnapshots:        limits.MaxTotalSnapshots,
			MaxTotalVolumeGigabytes:  limits.MaxTotalVolumeGigabytes,
			MaxTotalBackups:          limits.MaxTotalBackups,
			MaxTotalBackupGigabytes:  limits.MaxTotalBackupGigabytes,
			TotalVolumesUsed:         limits.TotalVolumesUsed,
			TotalGigabytesUsed:       limits.TotalGigabytesUsed,
			TotalSnapshotsUsed:       limits.TotalSnapshotsUsed,
			TotalBackupsUsed:         limits.TotalBackupsUsed,
			TotalBackupGigabytesUsed: limits.TotalBackupGigabytesUsed,
		},
		Rate: []rateLimitDocument{},
	}
}
