package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type createServerGroupRequest struct {
	ServerGroup createServerGroupDocument `json:"server_group"`
}

type createServerGroupDocument struct {
	Name     string   `json:"name"`
	Policies []string `json:"policies"`
	Policy   string   `json:"policy"`
}

type serverGroupListResponse struct {
	ServerGroups []serverGroupDocument `json:"server_groups"`
}

type serverGroupResponse struct {
	ServerGroup serverGroupDocument `json:"server_group"`
}

type serverGroupDocument struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Policies  []string       `json:"policies"`
	Members   []string       `json:"members"`
	UserID    string         `json:"user_id"`
	ProjectID string         `json:"project_id"`
	Metadata  map[string]any `json:"metadata"`
	Policy    *string        `json:"policy,omitempty"`
}

func (request createServerGroupRequest) createServerGroup() appcompute.CreateServerGroup {
	return appcompute.CreateServerGroup{
		Name:     request.ServerGroup.Name,
		Policies: request.ServerGroup.Policies,
		Policy:   request.ServerGroup.Policy,
	}
}

func toServerGroupDocuments(
	groups []appcompute.ServerGroup,
) []serverGroupDocument {
	documents := make([]serverGroupDocument, 0, len(groups))
	for _, group := range groups {
		documents = append(documents, toServerGroupDocument(group))
	}

	return documents
}

func toServerGroupDocument(group appcompute.ServerGroup) serverGroupDocument {
	var policy *string
	if group.Policy != "" {
		policy = &group.Policy
	}

	return serverGroupDocument{
		ID:        group.ID,
		Name:      group.Name,
		Policies:  group.Policies,
		Members:   group.Members,
		UserID:    group.UserID,
		ProjectID: group.ProjectID,
		Metadata:  group.Metadata,
		Policy:    policy,
	}
}
