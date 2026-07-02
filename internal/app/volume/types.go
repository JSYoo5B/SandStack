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

type AvailabilityZone struct {
	Name      string
	Available bool
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

type Transfer struct {
	ID        string
	AuthKey   string
	Name      string
	VolumeID  string
	CreatedAt string
	Links     []map[string]string
}

type CreateTransfer struct {
	Name     string
	VolumeID string
}

type Backup struct {
	ID                  string
	Name                string
	Description         string
	VolumeID            string
	SnapshotID          string
	Status              string
	Size                int
	ObjectCount         int
	Container           string
	HasDependentBackups bool
	FailReason          string
	IsIncremental       bool
	ProjectID           string
	Metadata            map[string]string
	AvailabilityZone    string
	CreatedAt           string
	UpdatedAt           string
	DataTimestamp       string
}

type CreateBackup struct {
	Name             string
	Description      string
	VolumeID         string
	Force            bool
	Metadata         map[string]string
	Container        string
	Incremental      bool
	SnapshotID       string
	AvailabilityZone string
}

type Attachment struct {
	ID             string
	VolumeID       string
	Instance       string
	AttachedAt     string
	DetachedAt     string
	Status         string
	AttachMode     string
	ConnectionInfo map[string]any
	Connector      map[string]any
}

type CreateAttachment struct {
	VolumeID   string
	InstanceID string
	Connector  map[string]any
	Mode       string
}

type UpdateAttachment struct {
	Connector map[string]any
}
