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
	for _, service := range s.defaultServices() {
		s.repositories.Services.Save(service)
	}
	for _, endpoint := range s.defaultEndpoints() {
		s.repositories.Endpoints.Save(endpoint)
	}
}

func (s Service) Reset() {
	s.repositories.Projects.Reset()
	s.repositories.Users.Reset()
	s.repositories.Roles.Reset()
	s.repositories.Tokens.Reset()
	s.repositories.Services.Reset()
	s.repositories.Endpoints.Reset()
	s.SeedDefaults()
}
