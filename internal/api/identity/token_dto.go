package identity

import (
	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
)

// authRequest models the Keystone wire JSON accepted by SandStack.
// Gophercloud's auth builders are client-side request builders, so the server
// keeps a small DTO for decoding inbound JSON and verifies compatibility with
// Gophercloud integration tests.
type authRequest struct {
	Auth struct {
		Identity struct {
			Methods  []string `json:"methods"`
			Password struct {
				User struct {
					Name     string `json:"name"`
					Password string `json:"password"`
					Domain   struct {
						ID string `json:"id"`
					} `json:"domain"`
				} `json:"user"`
			} `json:"password"`
		} `json:"identity"`
		Scope struct {
			Project struct {
				Name   string `json:"name"`
				Domain struct {
					ID string `json:"id"`
				} `json:"domain"`
			} `json:"project"`
		} `json:"scope"`
	} `json:"auth"`
}

func (r authRequest) passwordAuth() appidentity.PasswordAuth {
	return appidentity.PasswordAuth{
		Username:    r.Auth.Identity.Password.User.Name,
		Password:    r.Auth.Identity.Password.User.Password,
		ProjectName: r.Auth.Scope.Project.Name,
	}
}

type tokenResponse struct {
	Token issuedToken `json:"token"`
}

type issuedToken struct {
	ExpiresAt string                `json:"expires_at"`
	IssuedAt  string                `json:"issued_at"`
	Methods   []string              `json:"methods"`
	User      tokens.User           `json:"user"`
	Project   tokens.Project        `json:"project"`
	Roles     []tokens.Role         `json:"roles"`
	Catalog   []tokens.CatalogEntry `json:"catalog"`
}

func toIssuedToken(token appidentity.IssuedToken) issuedToken {
	return issuedToken{
		ExpiresAt: token.ExpiresAt,
		IssuedAt:  token.IssuedAt,
		Methods:   token.Methods,
		User:      token.User,
		Project:   token.Project,
		Roles:     token.Roles,
		Catalog:   token.Catalog,
	}
}
