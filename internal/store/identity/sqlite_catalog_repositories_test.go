package identity_test

import (
	"path/filepath"
	"testing"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	storeidentity "github.com/JSYoo5B/SandStack/internal/store/identity"
	"github.com/stretchr/testify/suite"
)

type SQLiteCatalogRepositoriesSuite struct {
	suite.Suite
	repositories *storeidentity.SQLiteCatalogRepositories
}

func TestSQLiteCatalogRepositoriesSuite(t *testing.T) {
	suite.Run(t, new(SQLiteCatalogRepositoriesSuite))
}

func (s *SQLiteCatalogRepositoriesSuite) SetupTest() {
	repositories, err := storeidentity.OpenSQLiteCatalogRepositories(":memory:")
	s.Require().NoError(err)

	s.repositories = repositories
}

func (s *SQLiteCatalogRepositoriesSuite) TearDownTest() {
	s.Require().NoError(s.repositories.Close())
}

func (s *SQLiteCatalogRepositoriesSuite) TestCreateListAndGetService() {
	created := s.repositories.Services.Save(testServiceDefinition())

	listed := s.repositories.Services.List()
	found, err := s.repositories.Services.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Len(listed, 1)
	s.Assert().Equal(created, listed[0])
	s.Assert().Equal(created, found)
}

func (s *SQLiteCatalogRepositoriesSuite) TestCreateListAndGetEndpoint() {
	created := s.repositories.Endpoints.Save(testEndpointDefinition())

	listed := s.repositories.Endpoints.List()
	byService := s.repositories.Endpoints.ListByServiceID("identity")
	found, err := s.repositories.Endpoints.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Len(listed, 1)
	s.Assert().Equal(created, listed[0])
	s.Assert().Equal([]appidentity.EndpointDefinition{created}, byService)
	s.Assert().Equal(created, found)
}

func (s *SQLiteCatalogRepositoriesSuite) TestResetClearsCatalog() {
	s.repositories.Services.Save(testServiceDefinition())
	s.repositories.Endpoints.Save(testEndpointDefinition())

	s.repositories.Endpoints.Reset()
	s.repositories.Services.Reset()

	s.Assert().Empty(s.repositories.Endpoints.List())
	s.Assert().Empty(s.repositories.Services.List())
}

func (s *SQLiteCatalogRepositoriesSuite) TestFileBackedDatabasePersistsCatalog() {
	path := filepath.Join(s.T().TempDir(), "sandstack.db")
	repositories, err := storeidentity.OpenSQLiteCatalogRepositories(path)
	s.Require().NoError(err)

	service := repositories.Services.Save(testServiceDefinition())
	endpoint := repositories.Endpoints.Save(testEndpointDefinition())
	s.Require().NoError(repositories.Close())

	reopened, err := storeidentity.OpenSQLiteCatalogRepositories(path)
	s.Require().NoError(err)
	defer reopened.Close()

	foundService, err := reopened.Services.Get(service.ID)
	s.Require().NoError(err)
	foundEndpoint, err := reopened.Endpoints.Get(endpoint.ID)
	s.Require().NoError(err)

	s.Assert().Equal(service, foundService)
	s.Assert().Equal(endpoint, foundEndpoint)
}

func testServiceDefinition() appidentity.ServiceDefinition {
	return appidentity.ServiceDefinition{
		ID:          "identity",
		Name:        "identity",
		Type:        "identity",
		Description: "SandStack identity service",
		Enabled:     true,
	}
}

func testEndpointDefinition() appidentity.EndpointDefinition {
	return appidentity.EndpointDefinition{
		ID:          "identity-public",
		ServiceID:   "identity",
		Interface:   "public",
		Region:      "RegionOne",
		Path:        "/identity/v3",
		Enabled:     true,
		Description: "SandStack identity endpoint",
	}
}
