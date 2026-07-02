package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"

func (s Service) Roles() []roles.Role {
	return s.repositories.Roles.List()
}

func (s Service) defaultRole() roles.Role {
	return roles.Role{
		DomainID:    "default",
		ID:          s.config.RoleName,
		Name:        s.config.RoleName,
		Description: "Default SandStack role",
		Links:       map[string]any{},
		Extra:       map[string]any{},
		Options:     map[roles.Option]any{},
	}
}

func (s Service) RoleByID(id string) (roles.Role, bool) {
	role, err := s.repositories.Roles.Get(id)
	if err != nil {
		return roles.Role{}, false
	}

	return role, true
}
