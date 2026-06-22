package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type portListResponse struct {
	Ports []portDocument `json:"ports"`
}

type portDocument struct {
	ID        string `json:"id"`
	NetworkID string `json:"network_id"`
	Name      string `json:"name"`
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
		ID:        port.ID,
		NetworkID: port.NetworkID,
		Name:      port.Name,
	}
}
