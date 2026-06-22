package image

type CreateImage struct {
	Name            string
	ContainerFormat string
	DiskFormat      string
	MinDisk         int
	MinRAM          int
	Tags            []string
}

type Image struct {
	ID              string
	Name            string
	Status          string
	ContainerFormat string
	DiskFormat      string
	MinDisk         int
	MinRAM          int
	Protected       bool
	Visibility      string
	Tags            []string
	CreatedAt       string
	UpdatedAt       string
}
