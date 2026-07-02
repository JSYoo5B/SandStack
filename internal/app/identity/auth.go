package identity

import (
	"errors"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
)

var ErrInvalidCredentials = errors.New("invalid identity credentials")

func (s Service) AuthenticatePassword(auth PasswordAuth, baseURL string) (IssuedToken, error) {
	user, err := s.repositories.Users.FindByName(auth.Username)
	if err != nil || user.Password != auth.Password {
		return IssuedToken{}, ErrInvalidCredentials
	}

	project, err := s.repositories.Projects.FindByName(auth.ProjectName)
	if err != nil {
		return IssuedToken{}, ErrInvalidCredentials
	}

	now := time.Now().UTC()
	roles := s.repositories.Roles.List()
	issuedToken := IssuedToken{
		ID:        "sandstack-" + idgen.RandomHex(16),
		ExpiresAt: now.Add(24 * time.Hour).Format(time.RFC3339),
		IssuedAt:  now.Format(time.RFC3339),
		Methods:   []string{"password"},
		User: tokens.User{
			ID:     user.ID,
			Name:   user.Name,
			Domain: defaultDomain(),
		},
		Project: tokens.Project{
			ID:     project.ID,
			Name:   project.Name,
			Domain: defaultDomain(),
		},
		Roles:   tokenRoles(roles),
		Catalog: s.Catalog(baseURL),
	}

	return s.repositories.Tokens.Save(issuedToken), nil
}

func (s Service) TokenByID(id string) (IssuedToken, error) {
	return s.repositories.Tokens.Get(id)
}

func (s Service) RevokeToken(id string) error {
	return s.repositories.Tokens.Delete(id)
}

func tokenRoles(roles []roles.Role) []tokens.Role {
	result := make([]tokens.Role, 0, len(roles))
	for _, role := range roles {
		result = append(result, tokens.Role{
			ID:   role.ID,
			Name: role.Name,
		})
	}

	return result
}

func defaultDomain() tokens.Domain {
	return tokens.Domain{
		ID:   "default",
		Name: "Default",
	}
}
