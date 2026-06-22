package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type createSubnetRequest struct {
	Subnet createSubnetDocument `json:"subnet"`
}

type createSubnetDocument struct {
	NetworkID      string   `json:"network_id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	IPVersion      int      `json:"ip_version"`
	CIDR           string   `json:"cidr"`
	GatewayIP      *string  `json:"gateway_ip"`
	DNSNameservers []string `json:"dns_nameservers"`
	EnableDHCP     *bool    `json:"enable_dhcp"`
	ProjectID      string   `json:"project_id"`
	TenantID       string   `json:"tenant_id"`
}

func (r createSubnetRequest) createSubnet() appnetwork.CreateSubnet {
	projectID := r.Subnet.ProjectID
	if projectID == "" {
		projectID = r.Subnet.TenantID
	}

	gatewayIP := ""
	if r.Subnet.GatewayIP != nil {
		gatewayIP = *r.Subnet.GatewayIP
	}

	return appnetwork.CreateSubnet{
		NetworkID:      r.Subnet.NetworkID,
		Name:           r.Subnet.Name,
		Description:    r.Subnet.Description,
		IPVersion:      r.Subnet.IPVersion,
		CIDR:           r.Subnet.CIDR,
		GatewayIP:      gatewayIP,
		DNSNameservers: r.Subnet.DNSNameservers,
		EnableDHCP:     r.Subnet.EnableDHCP,
		ProjectID:      projectID,
	}
}

type subnetListResponse struct {
	Subnets []subnetDocument `json:"subnets"`
}

type subnetResponse struct {
	Subnet subnetDocument `json:"subnet"`
}

type subnetDocument struct {
	ID              string                 `json:"id"`
	NetworkID       string                 `json:"network_id"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	IPVersion       int                    `json:"ip_version"`
	CIDR            string                 `json:"cidr"`
	GatewayIP       string                 `json:"gateway_ip"`
	DNSNameservers  []string               `json:"dns_nameservers"`
	AllocationPools []allocationPoolObject `json:"allocation_pools"`
	HostRoutes      []hostRouteObject      `json:"host_routes"`
	EnableDHCP      bool                   `json:"enable_dhcp"`
	TenantID        string                 `json:"tenant_id"`
	ProjectID       string                 `json:"project_id"`
}

type allocationPoolObject struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type hostRouteObject struct {
	DestinationCIDR string `json:"destination"`
	NextHop         string `json:"nexthop"`
}

func toSubnetDocuments(subnets []appnetwork.Subnet) []subnetDocument {
	documents := make([]subnetDocument, 0, len(subnets))
	for _, subnet := range subnets {
		documents = append(documents, toSubnetDocument(subnet))
	}

	return documents
}

func toSubnetDocument(subnet appnetwork.Subnet) subnetDocument {
	return subnetDocument{
		ID:              subnet.ID,
		NetworkID:       subnet.NetworkID,
		Name:            subnet.Name,
		Description:     subnet.Description,
		IPVersion:       subnet.IPVersion,
		CIDR:            subnet.CIDR,
		GatewayIP:       subnet.GatewayIP,
		DNSNameservers:  subnet.DNSNameservers,
		AllocationPools: toAllocationPoolObjects(subnet.AllocationPools),
		HostRoutes:      toHostRouteObjects(subnet.HostRoutes),
		EnableDHCP:      subnet.EnableDHCP,
		TenantID:        subnet.TenantID,
		ProjectID:       subnet.ProjectID,
	}
}

func toAllocationPoolObjects(
	pools []appnetwork.AllocationPool,
) []allocationPoolObject {
	objects := make([]allocationPoolObject, 0, len(pools))
	for _, pool := range pools {
		objects = append(objects, allocationPoolObject{
			Start: pool.Start,
			End:   pool.End,
		})
	}

	return objects
}

func toHostRouteObjects(routes []appnetwork.HostRoute) []hostRouteObject {
	objects := make([]hostRouteObject, 0, len(routes))
	for _, route := range routes {
		objects = append(objects, hostRouteObject{
			DestinationCIDR: route.DestinationCIDR,
			NextHop:         route.NextHop,
		})
	}

	return objects
}
