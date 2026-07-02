package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type createFloatingIPRequest struct {
	FloatingIP createFloatingIPDocument `json:"floatingip"`
}

type createFloatingIPDocument struct {
	Description       string `json:"description"`
	FloatingNetworkID string `json:"floating_network_id"`
	FloatingIP        string `json:"floating_ip_address"`
	PortID            string `json:"port_id"`
	FixedIP           string `json:"fixed_ip_address"`
	SubnetID          string `json:"subnet_id"`
	ProjectID         string `json:"project_id"`
	TenantID          string `json:"tenant_id"`
}

func (r createFloatingIPRequest) createFloatingIP() appnetwork.CreateFloatingIP {
	projectID := r.FloatingIP.ProjectID
	if projectID == "" {
		projectID = r.FloatingIP.TenantID
	}

	return appnetwork.CreateFloatingIP{
		Description:       r.FloatingIP.Description,
		FloatingNetworkID: r.FloatingIP.FloatingNetworkID,
		FloatingIP:        r.FloatingIP.FloatingIP,
		PortID:            r.FloatingIP.PortID,
		FixedIP:           r.FloatingIP.FixedIP,
		SubnetID:          r.FloatingIP.SubnetID,
		ProjectID:         projectID,
	}
}

type floatingIPListResponse struct {
	FloatingIPs []floatingIPDocument `json:"floatingips"`
}

type floatingIPResponse struct {
	FloatingIP floatingIPDocument `json:"floatingip"`
}

type floatingIPDocument struct {
	ID                string   `json:"id"`
	Description       string   `json:"description"`
	FloatingNetworkID string   `json:"floating_network_id"`
	FloatingIP        string   `json:"floating_ip_address"`
	PortID            string   `json:"port_id"`
	FixedIP           string   `json:"fixed_ip_address"`
	TenantID          string   `json:"tenant_id"`
	ProjectID         string   `json:"project_id"`
	Status            string   `json:"status"`
	RouterID          string   `json:"router_id"`
	Tags              []string `json:"tags"`
}

func toFloatingIPDocuments(
	floatingIPs []appnetwork.FloatingIP,
) []floatingIPDocument {
	documents := make([]floatingIPDocument, 0, len(floatingIPs))
	for _, floatingIP := range floatingIPs {
		documents = append(documents, toFloatingIPDocument(floatingIP))
	}

	return documents
}

func toFloatingIPDocument(floatingIP appnetwork.FloatingIP) floatingIPDocument {
	return floatingIPDocument{
		ID:                floatingIP.ID,
		Description:       floatingIP.Description,
		FloatingNetworkID: floatingIP.FloatingNetworkID,
		FloatingIP:        floatingIP.FloatingIP,
		PortID:            floatingIP.PortID,
		FixedIP:           floatingIP.FixedIP,
		TenantID:          floatingIP.TenantID,
		ProjectID:         floatingIP.ProjectID,
		Status:            floatingIP.Status,
		RouterID:          floatingIP.RouterID,
		Tags:              floatingIP.Tags,
	}
}
