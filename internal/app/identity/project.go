package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"

func (s Service) Projects() []projects.Project {
	return []projects.Project{s.Project()}
}

func (s Service) Project() projects.Project {
	return projects.Project{
		ID:          s.config.ProjectID,
		Name:        s.config.ProjectName,
		Description: "Default SandStack project",
		DomainID:    "default",
		Enabled:     true,
		IsDomain:    false,
		Tags:        []string{},
		Extra:       map[string]any{},
	}
}

func (s Service) ProjectByID(id string) (projects.Project, bool) {
	project := s.Project()
	if project.ID != id {
		return projects.Project{}, false
	}

	return project, true
}
