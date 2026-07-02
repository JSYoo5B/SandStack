package identity

import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"

func (s Service) Projects() []projects.Project {
	return s.repositories.Projects.List()
}

func (s Service) defaultProject() projects.Project {
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
	project, err := s.repositories.Projects.Get(id)
	if err != nil {
		return projects.Project{}, false
	}

	return project, true
}
