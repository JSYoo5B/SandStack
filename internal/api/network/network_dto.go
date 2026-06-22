package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type createNetworkRequest struct {
	Network struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		AdminStateUp *bool  `json:"admin_state_up"`
		ProjectID    string `json:"project_id"`
		TenantID     string `json:"tenant_id"`
		Shared       bool   `json:"shared"`
	} `json:"network"`
}

func (r createNetworkRequest) createNetwork() appnetwork.CreateNetwork {
	projectID := r.Network.ProjectID
	if projectID == "" {
		projectID = r.Network.TenantID
	}

	return appnetwork.CreateNetwork{
		Name:         r.Network.Name,
		Description:  r.Network.Description,
		AdminStateUp: r.Network.AdminStateUp,
		ProjectID:    projectID,
		Shared:       r.Network.Shared,
	}
}

type networkListResponse struct {
	Networks []networkDocument `json:"networks"`
}

type networkResponse struct {
	Network networkDocument `json:"network"`
}

type networkDocument struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	AdminStateUp bool     `json:"admin_state_up"`
	Status       string   `json:"status"`
	Subnets      []string `json:"subnets"`
	TenantID     string   `json:"tenant_id"`
	ProjectID    string   `json:"project_id"`
	Shared       bool     `json:"shared"`
}

func toNetworkDocuments(networks []appnetwork.Network) []networkDocument {
	documents := make([]networkDocument, 0, len(networks))
	for _, network := range networks {
		documents = append(documents, toNetworkDocument(network))
	}

	return documents
}

func toNetworkDocument(network appnetwork.Network) networkDocument {
	return networkDocument{
		ID:           network.ID,
		Name:         network.Name,
		Description:  network.Description,
		AdminStateUp: network.AdminStateUp,
		Status:       network.Status,
		Subnets:      network.Subnets,
		TenantID:     network.TenantID,
		ProjectID:    network.ProjectID,
		Shared:       network.Shared,
	}
}
