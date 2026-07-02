package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/endpoints"

type endpointsResponse struct {
	Endpoints []endpoints.Endpoint `json:"endpoints"`
	Links     identityLinks        `json:"links"`
}

type endpointResponse struct {
	Endpoint endpoints.Endpoint `json:"endpoint"`
}
