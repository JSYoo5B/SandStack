package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"

type User struct {
	DefaultProjectID string
	Description      string
	DomainID         string
	Enabled          bool
	ID               string
	Name             string
	Password         string
}

type ServiceDefinition struct {
	ID          string
	Name        string
	Type        string
	Description string
	Enabled     bool
}

type EndpointDefinition struct {
	ID          string
	ServiceID   string
	Interface   string
	Region      string
	Path        string
	Enabled     bool
	Description string
}

type PasswordAuth struct {
	Username    string
	Password    string
	ProjectName string
}

type IssuedToken struct {
	ID        string
	ExpiresAt string
	IssuedAt  string
	Methods   []string
	User      tokens.User
	Project   tokens.Project
	Roles     []tokens.Role
	Catalog   []tokens.CatalogEntry
}

type VersionDocument struct {
	ID         string
	Status     string
	Updated    string
	Links      []VersionLink
	MediaTypes []MediaType
}

type VersionLink struct {
	Rel  string
	Href string
}

type MediaType struct {
	Base string
	Type string
}
