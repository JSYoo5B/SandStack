package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type createPortRequest struct {
	Port createPortDocument `json:"port"`
}

type createPortDocument struct {
	NetworkID    string            `json:"network_id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	AdminStateUp *bool             `json:"admin_state_up"`
	FixedIPs     []fixedIPDocument `json:"fixed_ips"`
	ProjectID    string            `json:"project_id"`
	TenantID     string            `json:"tenant_id"`
	DeviceID     string            `json:"device_id"`
	DeviceOwner  string            `json:"device_owner"`
}

func (r createPortRequest) createPort() appnetwork.CreatePort {
	projectID := r.Port.ProjectID
	if projectID == "" {
		projectID = r.Port.TenantID
	}

	return appnetwork.CreatePort{
		NetworkID:    r.Port.NetworkID,
		Name:         r.Port.Name,
		Description:  r.Port.Description,
		AdminStateUp: r.Port.AdminStateUp,
		FixedIPs:     toAppFixedIPs(r.Port.FixedIPs),
		ProjectID:    projectID,
		DeviceID:     r.Port.DeviceID,
		DeviceOwner:  r.Port.DeviceOwner,
	}
}

type portListResponse struct {
	Ports []portDocument `json:"ports"`
}

type portResponse struct {
	Port portDocument `json:"port"`
}

type portDocument struct {
	ID           string            `json:"id"`
	NetworkID    string            `json:"network_id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	AdminStateUp bool              `json:"admin_state_up"`
	Status       string            `json:"status"`
	MACAddress   string            `json:"mac_address"`
	FixedIPs     []fixedIPDocument `json:"fixed_ips"`
	TenantID     string            `json:"tenant_id"`
	ProjectID    string            `json:"project_id"`
	DeviceID     string            `json:"device_id"`
	DeviceOwner  string            `json:"device_owner"`
}

type fixedIPDocument struct {
	SubnetID  string `json:"subnet_id,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
}

func toPortDocuments(ports []appnetwork.Port) []portDocument {
	documents := make([]portDocument, 0, len(ports))
	for _, port := range ports {
		documents = append(documents, toPortDocument(port))
	}

	return documents
}

func toPortDocument(port appnetwork.Port) portDocument {
	return portDocument{
		ID:           port.ID,
		NetworkID:    port.NetworkID,
		Name:         port.Name,
		Description:  port.Description,
		AdminStateUp: port.AdminStateUp,
		Status:       port.Status,
		MACAddress:   port.MACAddress,
		FixedIPs:     toFixedIPDocuments(port.FixedIPs),
		TenantID:     port.TenantID,
		ProjectID:    port.ProjectID,
		DeviceID:     port.DeviceID,
		DeviceOwner:  port.DeviceOwner,
	}
}

func toAppFixedIPs(fixedIPs []fixedIPDocument) []appnetwork.FixedIP {
	values := make([]appnetwork.FixedIP, 0, len(fixedIPs))
	for _, fixedIP := range fixedIPs {
		values = append(values, appnetwork.FixedIP{
			SubnetID:  fixedIP.SubnetID,
			IPAddress: fixedIP.IPAddress,
		})
	}

	return values
}

func toFixedIPDocuments(fixedIPs []appnetwork.FixedIP) []fixedIPDocument {
	documents := make([]fixedIPDocument, 0, len(fixedIPs))
	for _, fixedIP := range fixedIPs {
		documents = append(documents, fixedIPDocument{
			SubnetID:  fixedIP.SubnetID,
			IPAddress: fixedIP.IPAddress,
		})
	}

	return documents
}
