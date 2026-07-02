package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/services"

type servicesResponse struct {
	Services []services.Service `json:"services"`
	Links    identityLinks      `json:"links"`
}

type serviceResponse struct {
	Service services.Service `json:"service"`
}
