package compute_test

import (
	"path/filepath"
	"testing"

	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	storecompute "github.com/JSYoo5B/SandStack/internal/store/compute"
	"github.com/stretchr/testify/suite"
)

type SQLiteServerRepositorySuite struct {
	suite.Suite
	repository *storecompute.SQLiteServerRepository
}

func TestSQLiteServerRepositorySuite(t *testing.T) {
	suite.Run(t, new(SQLiteServerRepositorySuite))
}

func (s *SQLiteServerRepositorySuite) SetupTest() {
	repository, err := storecompute.OpenSQLiteServerRepository(":memory:")
	s.Require().NoError(err)

	s.repository = repository
}

func (s *SQLiteServerRepositorySuite) TearDownTest() {
	s.Require().NoError(s.repository.Close())
}

func (s *SQLiteServerRepositorySuite) TestCreateListGetAndUpdateServer() {
	created := s.repository.Create(serverFixture("srv-1"))
	created.Status = "ACTIVE"
	created.Progress = 100

	updated, err := s.repository.Update(created)
	s.Require().NoError(err)
	listed := s.repository.List()
	found, err := s.repository.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Len(listed, 1)
	s.Assert().Equal(updated, listed[0])
	s.Assert().Equal(updated, found)
}

func (s *SQLiteServerRepositorySuite) TestDeleteServer() {
	created := s.repository.Create(serverFixture("srv-1"))

	err := s.repository.Delete(created.ID)
	s.Require().NoError(err)

	_, err = s.repository.Get(created.ID)
	s.Require().ErrorIs(err, appcompute.ErrServerNotFound)
	s.Assert().Empty(s.repository.List())
}

func (s *SQLiteServerRepositorySuite) TestResetClearsServers() {
	s.repository.Create(serverFixture("srv-1"))

	s.repository.Reset()

	s.Assert().Empty(s.repository.List())
}

func (s *SQLiteServerRepositorySuite) TestFileBackedDatabasePersistsServers() {
	path := filepath.Join(s.T().TempDir(), "sandstack.db")
	repository, err := storecompute.OpenSQLiteServerRepository(path)
	s.Require().NoError(err)

	created := repository.Create(serverFixture("srv-1"))
	s.Require().NoError(repository.Close())

	reopened, err := storecompute.OpenSQLiteServerRepository(path)
	s.Require().NoError(err)
	defer reopened.Close()

	found, err := reopened.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Equal(created, found)
}

func serverFixture(id string) appcompute.Server {
	return appcompute.Server{
		ID:        id,
		Name:      "web",
		ImageID:   "img-1",
		FlavorID:  "1",
		TenantID:  "demo",
		UserID:    "admin",
		Status:    "BUILD",
		Progress:  0,
		CreatedAt: "2026-07-03T00:00:00Z",
		UpdatedAt: "2026-07-03T00:00:00Z",
		Metadata: map[string]string{
			"role": "web",
		},
	}
}
