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

type Subnet struct {
	ID              string
	NetworkID       string
	Name            string
	Description     string
	IPVersion       int
	CIDR            string
	GatewayIP       string
	DNSNameservers  []string
	AllocationPools []AllocationPool
	HostRoutes      []HostRoute
	EnableDHCP      bool
	TenantID        string
	ProjectID       string
}

type CreateSubnet struct {
	NetworkID      string
	Name           string
	Description    string
	IPVersion      int
	CIDR           string
	GatewayIP      string
	DNSNameservers []string
	EnableDHCP     *bool
	ProjectID      string
}

type AllocationPool struct {
	Start string
	End   string
}

type HostRoute struct {
	DestinationCIDR string
	NextHop         string
}

type Port struct {
	ID        string
	NetworkID string
	Name      string
}
