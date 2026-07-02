package identity

import "github.com/JSYoo5B/SandStack/internal/platform/config"

type Service struct {
	config       config.Config
	repositories Repositories
}

func NewServiceWithRepositories(
	cfg config.Config,
	repositories Repositories,
) Service {
	service := Service{
		config:       cfg,
		repositories: repositories,
	}
	service.SeedDefaults()

	return service
}

func (s Service) SeedDefaults() {
	s.repositories.Projects.Save(s.defaultProject())
	s.repositories.Users.Save(s.defaultUser())
	s.repositories.Roles.Save(s.defaultRole())
}

func (s Service) Reset() {
	s.repositories.Projects.Reset()
	s.repositories.Users.Reset()
	s.repositories.Roles.Reset()
	s.SeedDefaults()
}
