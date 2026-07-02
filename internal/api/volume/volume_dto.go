package volume

import appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"

type createVolumeRequest struct {
	Volume struct {
		Size        int               `json:"size"`
		Name        string            `json:"name"`
		Description string            `json:"description"`
		VolumeType  string            `json:"volume_type"`
		Metadata    map[string]string `json:"metadata"`
	} `json:"volume"`
}

func (r createVolumeRequest) createVolume() appvolume.CreateVolume {
	return appvolume.CreateVolume{
		Size:        r.Volume.Size,
		Name:        r.Volume.Name,
		Description: r.Volume.Description,
		VolumeType:  r.Volume.VolumeType,
		Metadata:    r.Volume.Metadata,
	}
}

type volumeListResponse struct {
	Volumes []volumeDocument `json:"volumes"`
}

type volumeResponse struct {
	Volume volumeDocument `json:"volume"`
}

type volumeDocument struct {
	ID          string            `json:"id"`
	Status      string            `json:"status"`
	Size        int               `json:"size"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	VolumeType  string            `json:"volume_type"`
	Metadata    map[string]string `json:"metadata"`
	Bootable    string            `json:"bootable"`
	Encrypted   bool              `json:"encrypted"`
	Multiattach bool              `json:"multiattach"`
}

func toVolumeDocuments(volumes []appvolume.Volume) []volumeDocument {
	documents := make([]volumeDocument, 0, len(volumes))
	for _, volume := range volumes {
		documents = append(documents, toVolumeDocument(volume))
	}

	return documents
}

func toVolumeDocument(volume appvolume.Volume) volumeDocument {
	return volumeDocument{
		ID:          volume.ID,
		Status:      volume.Status,
		Size:        volume.Size,
		CreatedAt:   volume.CreatedAt,
		UpdatedAt:   volume.UpdatedAt,
		Name:        volume.Name,
		Description: volume.Description,
		VolumeType:  volume.VolumeType,
		Metadata:    volume.Metadata,
		Bootable:    volume.Bootable,
		Encrypted:   volume.Encrypted,
		Multiattach: volume.Multiattach,
	}
}
