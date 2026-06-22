package volume

type CreateVolume struct {
	Size        int
	Name        string
	Description string
	VolumeType  string
	Metadata    map[string]string
}

type Volume struct {
	ID          string
	Status      string
	Size        int
	Name        string
	Description string
	VolumeType  string
	Metadata    map[string]string
	CreatedAt   string
	UpdatedAt   string
	Bootable    string
	Encrypted   bool
	Multiattach bool
}

type VolumeType struct {
	ID          string
	Name        string
	Description string
	ExtraSpecs  map[string]string
	IsPublic    bool
}
