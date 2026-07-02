package volume

import appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"

type volumeTypeListResponse struct {
	VolumeTypes []volumeTypeDocument `json:"volume_types"`
}

type volumeTypeResponse struct {
	VolumeType volumeTypeDocument `json:"volume_type"`
}

type volumeTypeDocument struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	ExtraSpecs   map[string]string `json:"extra_specs"`
	IsPublic     bool              `json:"is_public"`
	PublicAccess bool              `json:"os-volume-type-access:is_public"`
}

func toVolumeTypeDocuments(types []appvolume.VolumeType) []volumeTypeDocument {
	documents := make([]volumeTypeDocument, 0, len(types))
	for _, volumeType := range types {
		documents = append(documents, toVolumeTypeDocument(volumeType))
	}

	return documents
}

func toVolumeTypeDocument(volumeType appvolume.VolumeType) volumeTypeDocument {
	return volumeTypeDocument{
		ID:           volumeType.ID,
		Name:         volumeType.Name,
		Description:  volumeType.Description,
		ExtraSpecs:   volumeType.ExtraSpecs,
		IsPublic:     volumeType.IsPublic,
		PublicAccess: volumeType.IsPublic,
	}
}
