package volume

import appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"

type createBackupRequest struct {
	Backup createBackupDocument `json:"backup"`
}

type createBackupDocument struct {
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	VolumeID         string            `json:"volume_id"`
	Force            bool              `json:"force"`
	Metadata         map[string]string `json:"metadata"`
	Container        string            `json:"container"`
	Incremental      bool              `json:"incremental"`
	SnapshotID       string            `json:"snapshot_id"`
	AvailabilityZone string            `json:"availability_zone"`
}

func (r createBackupRequest) createBackup() appvolume.CreateBackup {
	return appvolume.CreateBackup{
		Name:             r.Backup.Name,
		Description:      r.Backup.Description,
		VolumeID:         r.Backup.VolumeID,
		Force:            r.Backup.Force,
		Metadata:         r.Backup.Metadata,
		Container:        r.Backup.Container,
		Incremental:      r.Backup.Incremental,
		SnapshotID:       r.Backup.SnapshotID,
		AvailabilityZone: r.Backup.AvailabilityZone,
	}
}

type backupListResponse struct {
	Backups []backupDocument `json:"backups"`
}

type backupResponse struct {
	Backup backupDocument `json:"backup"`
}

type backupDocument struct {
	ID                  string             `json:"id"`
	Name                string             `json:"name"`
	Description         string             `json:"description"`
	VolumeID            string             `json:"volume_id"`
	SnapshotID          string             `json:"snapshot_id"`
	Status              string             `json:"status"`
	Size                int                `json:"size"`
	ObjectCount         int                `json:"object_count"`
	Container           string             `json:"container"`
	HasDependentBackups bool               `json:"has_dependent_backups"`
	FailReason          string             `json:"fail_reason"`
	IsIncremental       bool               `json:"is_incremental"`
	ProjectID           string             `json:"os-backup-project-attr:project_id"`
	Metadata            *map[string]string `json:"metadata"`
	AvailabilityZone    *string            `json:"availability_zone"`
	CreatedAt           string             `json:"created_at"`
	UpdatedAt           string             `json:"updated_at"`
	DataTimestamp       string             `json:"data_timestamp"`
}

func toBackupDocuments(backups []appvolume.Backup) []backupDocument {
	documents := make([]backupDocument, 0, len(backups))
	for _, backup := range backups {
		documents = append(documents, toBackupDocument(backup))
	}

	return documents
}

func toBackupDocument(backup appvolume.Backup) backupDocument {
	metadata := backup.Metadata
	availabilityZone := backup.AvailabilityZone

	return backupDocument{
		ID:                  backup.ID,
		Name:                backup.Name,
		Description:         backup.Description,
		VolumeID:            backup.VolumeID,
		SnapshotID:          backup.SnapshotID,
		Status:              backup.Status,
		Size:                backup.Size,
		ObjectCount:         backup.ObjectCount,
		Container:           backup.Container,
		HasDependentBackups: backup.HasDependentBackups,
		FailReason:          backup.FailReason,
		IsIncremental:       backup.IsIncremental,
		ProjectID:           backup.ProjectID,
		Metadata:            &metadata,
		AvailabilityZone:    &availabilityZone,
		CreatedAt:           backup.CreatedAt,
		UpdatedAt:           backup.UpdatedAt,
		DataTimestamp:       backup.DataTimestamp,
	}
}
