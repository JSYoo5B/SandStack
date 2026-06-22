package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type createServerRequest struct {
	Server createServerDocument `json:"server"`
}

type createServerDocument struct {
	Name      string            `json:"name"`
	ImageRef  string            `json:"imageRef"`
	FlavorRef string            `json:"flavorRef"`
	Metadata  map[string]string `json:"metadata"`
}

func (r createServerRequest) createServer() appcompute.CreateServer {
	return appcompute.CreateServer{
		Name:     r.Server.Name,
		ImageID:  r.Server.ImageRef,
		FlavorID: r.Server.FlavorRef,
		Metadata: r.Server.Metadata,
	}
}

type serverListResponse struct {
	Servers []serverDocument `json:"servers"`
}

type serverResponse struct {
	Server serverDocument `json:"server"`
}

type serverDocument struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Image     map[string]any    `json:"image"`
	Flavor    map[string]any    `json:"flavor"`
	TenantID  string            `json:"tenant_id"`
	UserID    string            `json:"user_id"`
	Status    string            `json:"status"`
	Progress  int               `json:"progress"`
	CreatedAt string            `json:"created"`
	UpdatedAt string            `json:"updated"`
	Addresses map[string]any    `json:"addresses"`
	Metadata  map[string]string `json:"metadata"`
	Links     []map[string]any  `json:"links"`
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
		ID:        server.ID,
		Name:      server.Name,
		Image:     map[string]any{"id": server.ImageID},
		Flavor:    map[string]any{"id": server.FlavorID},
		TenantID:  server.TenantID,
		UserID:    server.UserID,
		Status:    server.Status,
		Progress:  server.Progress,
		CreatedAt: server.CreatedAt,
		UpdatedAt: server.UpdatedAt,
		Addresses: map[string]any{},
		Metadata:  server.Metadata,
		Links: []map[string]any{
			{
				"href": "/servers/" + server.ID,
				"rel":  "self",
			},
		},
	}
}
