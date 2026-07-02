package identity_test

import (
	"path/filepath"
	"testing"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	storeidentity "github.com/JSYoo5B/SandStack/internal/store/identity"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"
	"github.com/stretchr/testify/suite"
)

type SQLitePrincipalRepositoriesSuite struct {
	suite.Suite
	repositories *storeidentity.SQLitePrincipalRepositories
}

func TestSQLitePrincipalRepositoriesSuite(t *testing.T) {
	suite.Run(t, new(SQLitePrincipalRepositoriesSuite))
}

func (s *SQLitePrincipalRepositoriesSuite) SetupTest() {
	repositories, err := storeidentity.OpenSQLitePrincipalRepositories(":memory:")
	s.Require().NoError(err)

	s.repositories = repositories
}

func (s *SQLitePrincipalRepositoriesSuite) TearDownTest() {
	s.Require().NoError(s.repositories.Close())
}

func (s *SQLitePrincipalRepositoriesSuite) TestCreateListAndFindUser() {
	created := s.repositories.Users.Save(appidentity.User{
		ID:               "user-1",
		Name:             "admin",
		Password:         "password",
		DefaultProjectID: "demo",
		Description:      "Default user",
		DomainID:         "default",
		Enabled:          true,
	})

	listed := s.repositories.Users.List()
	found, err := s.repositories.Users.FindByName("admin")
	s.Require().NoError(err)

	s.Assert().Len(listed, 1)
	s.Assert().Equal(created, listed[0])
	s.Assert().Equal(created, found)
}

func (s *SQLitePrincipalRepositoriesSuite) TestCreateListAndFindProject() {
	created := s.repositories.Projects.Save(projects.Project{
		ID:          "demo",
		Name:        "demo",
		Description: "Default project",
		DomainID:    "default",
		Enabled:     true,
		Tags:        []string{},
		Extra:       map[string]any{},
	})

	listed := s.repositories.Projects.List()
	found, err := s.repositories.Projects.FindByName("demo")
	s.Require().NoError(err)

	s.Assert().Len(listed, 1)
	s.Assert().Equal(created, listed[0])
	s.Assert().Equal(created, found)
}

func (s *SQLitePrincipalRepositoriesSuite) TestCreateListAndGetRole() {
	created := s.repositories.Roles.Save(roles.Role{
		ID:          "admin",
		Name:        "admin",
		Description: "Default role",
		DomainID:    "default",
		Links:       map[string]any{},
		Extra:       map[string]any{},
		Options:     map[roles.Option]any{},
	})

	listed := s.repositories.Roles.List()
	found, err := s.repositories.Roles.Get("admin")
	s.Require().NoError(err)

	s.Assert().Len(listed, 1)
	s.Assert().Equal(created, listed[0])
	s.Assert().Equal(created, found)
}

func (s *SQLitePrincipalRepositoriesSuite) TestResetClearsPrincipals() {
	s.repositories.Users.Save(appidentity.User{
		ID:               "user-1",
		Name:             "admin",
		Password:         "password",
		DefaultProjectID: "demo",
		Description:      "Default user",
		DomainID:         "default",
		Enabled:          true,
	})

	s.repositories.Users.Reset()

	s.Assert().Empty(s.repositories.Users.List())
}

func (s *SQLitePrincipalRepositoriesSuite) TestFileBackedDatabasePersistsPrincipals() {
	path := filepath.Join(s.T().TempDir(), "sandstack.db")
	repositories, err := storeidentity.OpenSQLitePrincipalRepositories(path)
	s.Require().NoError(err)

	created := repositories.Users.Save(appidentity.User{
		ID:               "user-1",
		Name:             "admin",
		Password:         "password",
		DefaultProjectID: "demo",
		Description:      "Default user",
		DomainID:         "default",
		Enabled:          true,
	})
	s.Require().NoError(repositories.Close())

	reopened, err := storeidentity.OpenSQLitePrincipalRepositories(path)
	s.Require().NoError(err)
	defer reopened.Close()

	found, err := reopened.Users.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Equal(created, found)
}
