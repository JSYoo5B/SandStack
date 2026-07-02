package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"

type projectsResponse struct {
	Projects []projects.Project `json:"projects"`
	Links    identityLinks      `json:"links"`
}

type projectResponse struct {
	Project projects.Project `json:"project"`
}

type identityLinks struct {
	Self     string `json:"self"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
}
