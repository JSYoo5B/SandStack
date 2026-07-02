package identity

import (
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"
)

type Repositories struct {
	Users    UserRepository
	Projects ProjectRepository
	Roles    RoleRepository
}

type UserRepository interface {
	Save(user User) User
	List() []User
	Get(id string) (User, error)
	FindByName(name string) (User, error)
	Reset()
}

type ProjectRepository interface {
	Save(project projects.Project) projects.Project
	List() []projects.Project
	Get(id string) (projects.Project, error)
	FindByName(name string) (projects.Project, error)
	Reset()
}

type RoleRepository interface {
	Save(role roles.Role) roles.Role
	List() []roles.Role
	Get(id string) (roles.Role, error)
	Reset()
}
