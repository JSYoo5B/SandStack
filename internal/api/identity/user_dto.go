package identity

import appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"

type usersResponse struct {
	Users []user        `json:"users"`
	Links identityLinks `json:"links"`
}

type userResponse struct {
	User user `json:"user"`
}

type user struct {
	DefaultProjectID string         `json:"default_project_id"`
	Description      string         `json:"description"`
	DomainID         string         `json:"domain_id"`
	Enabled          bool           `json:"enabled"`
	ID               string         `json:"id"`
	Links            map[string]any `json:"links"`
	Name             string         `json:"name"`
	Options          map[string]any `json:"options"`
}

func toUser(source appidentity.User, baseURL string) user {
	return user{
		DefaultProjectID: source.DefaultProjectID,
		Description:      source.Description,
		DomainID:         source.DomainID,
		Enabled:          source.Enabled,
		ID:               source.ID,
		Links: map[string]any{
			"self": baseURL + "/identity/v3/users/" + source.ID,
		},
		Name:    source.Name,
		Options: map[string]any{},
	}
}
