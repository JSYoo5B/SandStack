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

type Snapshot struct {
	ID          string
	Name        string
	Description string
	VolumeID    string
	Status      string
	Size        int
	Metadata    map[string]string
	Progress    string
	ProjectID   string
	UserID      string
	CreatedAt   string
	UpdatedAt   string
}

type CreateSnapshot struct {
	Name        string
	Description string
	VolumeID    string
	Force       bool
	Metadata    map[string]string
}
