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
	ID           string
	NetworkID    string
	Name         string
	Description  string
	AdminStateUp bool
	Status       string
	MACAddress   string
	FixedIPs     []FixedIP
	TenantID     string
	ProjectID    string
	DeviceID     string
	DeviceOwner  string
}

type CreatePort struct {
	NetworkID    string
	Name         string
	Description  string
	AdminStateUp *bool
	FixedIPs     []FixedIP
	ProjectID    string
	DeviceID     string
	DeviceOwner  string
}

type FixedIP struct {
	SubnetID  string
	IPAddress string
}

type SecurityGroup struct {
	ID          string
	Name        string
	Description string
	Stateful    bool
	TenantID    string
	ProjectID   string
	Rules       []SecurityGroupRule
	Tags        []string
}

type CreateSecurityGroup struct {
	Name        string
	Description string
	Stateful    *bool
	ProjectID   string
}

type SecurityGroupRule struct {
	ID                   string
	Direction            string
	Description          string
	EtherType            string
	Protocol             string
	PortRangeMin         int
	PortRangeMax         int
	RemoteAddressGroupID string
	RemoteIPPrefix       string
	RemoteGroupID        string
	SecurityGroupID      string
	TenantID             string
	ProjectID            string
}

type CreateSecurityGroupRule struct {
	Direction            string
	Description          string
	EtherType            string
	Protocol             string
	PortRangeMin         int
	PortRangeMax         int
	RemoteAddressGroupID string
	RemoteIPPrefix       string
	RemoteGroupID        string
	SecurityGroupID      string
	ProjectID            string
}

type Router struct {
	ID                    string
	Name                  string
	Description           string
	AdminStateUp          bool
	Distributed           bool
	Status                string
	TenantID              string
	ProjectID             string
	GatewayInfo           RouterGatewayInfo
	Routes                []Route
	AvailabilityZoneHints []string
	Tags                  []string
}

type CreateRouter struct {
	Name                  string
	Description           string
	AdminStateUp          *bool
	Distributed           *bool
	ProjectID             string
	GatewayInfo           RouterGatewayInfo
	AvailabilityZoneHints []string
}

type RouterGatewayInfo struct {
	NetworkID        string
	EnableSNAT       *bool
	ExternalFixedIPs []ExternalFixedIP
	QoSPolicyID      string
}

type ExternalFixedIP struct {
	IPAddress string
	SubnetID  string
}

type Route struct {
	NextHop         string
	DestinationCIDR string
}
