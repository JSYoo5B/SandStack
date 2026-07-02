package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"

func (s Service) Roles() []roles.Role {
	return []roles.Role{s.Role()}
}

func (s Service) Role() roles.Role {
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
	role := s.Role()
	if role.ID != id {
		return roles.Role{}, false
	}

	return role, true
}
