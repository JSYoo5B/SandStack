package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type serverListResponse struct {
	Servers []serverDocument `json:"servers"`
}

type serverDocument struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func toServerDocuments(servers []appcompute.Server) []serverDocument {
	documents := make([]serverDocument, 0, len(servers))
	for _, server := range servers {
		documents = append(documents, toServerDocument(server))
	}

	return documents
}

func toServerDocument(server appcompute.Server) serverDocument {
	return serverDocument{
		ID:   server.ID,
		Name: server.Name,
	}
}
