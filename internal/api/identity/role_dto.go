package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"

type rolesResponse struct {
	Roles []roles.Role  `json:"roles"`
	Links identityLinks `json:"links"`
}

type roleResponse struct {
	Role roles.Role `json:"role"`
}
