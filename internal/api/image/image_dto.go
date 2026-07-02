package image

import appimage "github.com/JSYoo5B/SandStack/internal/app/image"

type createImageRequest struct {
	Name            string   `json:"name"`
	ContainerFormat string   `json:"container_format"`
	DiskFormat      string   `json:"disk_format"`
	MinDisk         int      `json:"min_disk"`
	MinRAM          int      `json:"min_ram"`
	Tags            []string `json:"tags"`
}

func (r createImageRequest) createImage() appimage.CreateImage {
	return appimage.CreateImage{
		Name:            r.Name,
		ContainerFormat: r.ContainerFormat,
		DiskFormat:      r.DiskFormat,
		MinDisk:         r.MinDisk,
		MinRAM:          r.MinRAM,
		Tags:            r.Tags,
	}
}

type imageListResponse struct {
	Images []imageDocument `json:"images"`
}

type imageDocument struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Status          string   `json:"status"`
	ContainerFormat string   `json:"container_format"`
	DiskFormat      string   `json:"disk_format"`
	MinDisk         int      `json:"min_disk"`
	MinRAM          int      `json:"min_ram"`
	Protected       bool     `json:"protected"`
	Visibility      string   `json:"visibility"`
	Tags            []string `json:"tags"`
	SizeBytes       int64    `json:"size"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	File            string   `json:"file"`
	Schema          string   `json:"schema"`
}

func toImageDocuments(images []appimage.Image) []imageDocument {
	documents := make([]imageDocument, 0, len(images))
	for _, image := range images {
		documents = append(documents, toImageDocument(image))
	}

	return documents
}

func toImageDocument(image appimage.Image) imageDocument {
	return imageDocument{
		ID:              image.ID,
		Name:            image.Name,
		Status:          image.Status,
		ContainerFormat: image.ContainerFormat,
		DiskFormat:      image.DiskFormat,
		MinDisk:         image.MinDisk,
		MinRAM:          image.MinRAM,
		Protected:       image.Protected,
		Visibility:      image.Visibility,
		Tags:            image.Tags,
		SizeBytes:       image.SizeBytes,
		CreatedAt:       image.CreatedAt,
		UpdatedAt:       image.UpdatedAt,
		File:            "/v2/images/" + image.ID + "/file",
		Schema:          "/v2/schemas/image",
	}
}
