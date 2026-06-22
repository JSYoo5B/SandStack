package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type subnetListResponse struct {
	Subnets []subnetDocument `json:"subnets"`
}

type subnetDocument struct {
	ID        string `json:"id"`
	NetworkID string `json:"network_id"`
	Name      string `json:"name"`
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
		ID:        subnet.ID,
		NetworkID: subnet.NetworkID,
		Name:      subnet.Name,
	}
}
