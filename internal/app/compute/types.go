package compute

type Flavor struct {
	ID          string
	Name        string
	RAM         int
	VCPUs       int
	Disk        int
	Swap        int
	RxTxFactor  float64
	IsPublic    bool
	Ephemeral   int
	Description string
	ExtraSpecs  map[string]string
}

type Server struct {
	ID        string
	Name      string
	ImageID   string
	FlavorID  string
	TenantID  string
	UserID    string
	Status    string
	Progress  int
	CreatedAt string
	UpdatedAt string
	Metadata  map[string]string
}

type CreateServer struct {
	Name     string
	ImageID  string
	FlavorID string
	Metadata map[string]string
}

type KeyPair struct {
	Name        string
	Fingerprint string
	PublicKey   string
	PrivateKey  string
	UserID      string
	Type        string
}

type CreateKeyPair struct {
	Name      string
	UserID    string
	Type      string
	PublicKey string
}
