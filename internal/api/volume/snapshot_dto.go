package volume

import appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"

type createSnapshotRequest struct {
	Snapshot createSnapshotDocument `json:"snapshot"`
}

type createSnapshotDocument struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	VolumeID    string            `json:"volume_id"`
	Force       bool              `json:"force"`
	Metadata    map[string]string `json:"metadata"`
}

func (r createSnapshotRequest) createSnapshot() appvolume.CreateSnapshot {
	return appvolume.CreateSnapshot{
		Name:        r.Snapshot.Name,
		Description: r.Snapshot.Description,
		VolumeID:    r.Snapshot.VolumeID,
		Force:       r.Snapshot.Force,
		Metadata:    r.Snapshot.Metadata,
	}
}

type snapshotListResponse struct {
	Snapshots []snapshotDocument `json:"snapshots"`
}

type snapshotResponse struct {
	Snapshot snapshotDocument `json:"snapshot"`
}

type snapshotDocument struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	VolumeID    string            `json:"volume_id"`
	Status      string            `json:"status"`
	Size        int               `json:"size"`
	Metadata    map[string]string `json:"metadata"`
	Progress    string            `json:"os-extended-snapshot-attributes:progress"`
	ProjectID   string            `json:"os-extended-snapshot-attributes:project_id"`
	UserID      string            `json:"user_id"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
}

func toSnapshotDocuments(snapshots []appvolume.Snapshot) []snapshotDocument {
	documents := make([]snapshotDocument, 0, len(snapshots))
	for _, snapshot := range snapshots {
		documents = append(documents, toSnapshotDocument(snapshot))
	}

	return documents
}

func toSnapshotDocument(snapshot appvolume.Snapshot) snapshotDocument {
	return snapshotDocument{
		ID:          snapshot.ID,
		Name:        snapshot.Name,
		Description: snapshot.Description,
		VolumeID:    snapshot.VolumeID,
		Status:      snapshot.Status,
		Size:        snapshot.Size,
		Metadata:    snapshot.Metadata,
		Progress:    snapshot.Progress,
		ProjectID:   snapshot.ProjectID,
		UserID:      snapshot.UserID,
		CreatedAt:   snapshot.CreatedAt,
		UpdatedAt:   snapshot.UpdatedAt,
	}
}
