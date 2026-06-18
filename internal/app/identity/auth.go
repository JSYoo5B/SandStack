package identity

import (
	"errors"
	"time"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
)

var ErrInvalidCredentials = errors.New("invalid identity credentials")

func (s Service) AuthenticatePassword(auth PasswordAuth, baseURL string) (IssuedToken, error) {
	if auth.Username != s.config.Username || auth.Password != s.config.Password || auth.ProjectName != s.config.ProjectName {
		return IssuedToken{}, ErrInvalidCredentials
	}

	now := time.Now().UTC()
	return IssuedToken{
		ID:        "sandstack-" + idgen.RandomHex(16),
		ExpiresAt: now.Add(24 * time.Hour).Format(time.RFC3339),
		IssuedAt:  now.Format(time.RFC3339),
		Methods:   []string{"password"},
		User: tokens.User{
			ID:     s.config.UserID,
			Name:   s.config.Username,
			Domain: defaultDomain(),
		},
		Project: tokens.Project{
			ID:     s.config.ProjectID,
			Name:   s.config.ProjectName,
			Domain: defaultDomain(),
		},
		Roles: []tokens.Role{
			{
				ID:   s.config.RoleName,
				Name: s.config.RoleName,
			},
		},
		Catalog: s.Catalog(baseURL),
	}, nil
}

func defaultDomain() tokens.Domain {
	return tokens.Domain{
		ID:   "default",
		Name: "Default",
	}
}
