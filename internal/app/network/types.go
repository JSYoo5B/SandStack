package network

type CreateNetwork struct {
	Name         string
	Description  string
	AdminStateUp *bool
	ProjectID    string
	Shared       bool
}

type Network struct {
	ID           string
	Name         string
	Description  string
	AdminStateUp bool
	Status       string
	Subnets      []string
	TenantID     string
	ProjectID    string
	Shared       bool
}
