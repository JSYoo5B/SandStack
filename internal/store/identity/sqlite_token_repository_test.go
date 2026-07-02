package identity_test

import (
	"path/filepath"
	"testing"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	storeidentity "github.com/JSYoo5B/SandStack/internal/store/identity"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
	"github.com/stretchr/testify/suite"
)

type SQLiteTokenRepositorySuite struct {
	suite.Suite
	repository *storeidentity.SQLiteTokenRepository
}

func TestSQLiteTokenRepositorySuite(t *testing.T) {
	suite.Run(t, new(SQLiteTokenRepositorySuite))
}

func (s *SQLiteTokenRepositorySuite) SetupTest() {
	repository, err := storeidentity.OpenSQLiteTokenRepository(":memory:")
	s.Require().NoError(err)

	s.repository = repository
}

func (s *SQLiteTokenRepositorySuite) TearDownTest() {
	s.Require().NoError(s.repository.Close())
}

func (s *SQLiteTokenRepositorySuite) TestSaveGetAndDeleteToken() {
	created := s.repository.Save(testIssuedToken())

	found, err := s.repository.Get(created.ID)
	s.Require().NoError(err)
	s.Assert().Equal(created, found)

	err = s.repository.Delete(created.ID)
	s.Require().NoError(err)

	_, err = s.repository.Get(created.ID)
	s.Require().ErrorIs(err, appidentity.ErrTokenNotFound)
}

func (s *SQLiteTokenRepositorySuite) TestResetClearsTokens() {
	s.repository.Save(testIssuedToken())

	s.repository.Reset()

	_, err := s.repository.Get("token-1")
	s.Require().ErrorIs(err, appidentity.ErrTokenNotFound)
}

func (s *SQLiteTokenRepositorySuite) TestFileBackedDatabasePersistsTokens() {
	path := filepath.Join(s.T().TempDir(), "sandstack.db")
	repository, err := storeidentity.OpenSQLiteTokenRepository(path)
	s.Require().NoError(err)

	created := repository.Save(testIssuedToken())
	s.Require().NoError(repository.Close())

	reopened, err := storeidentity.OpenSQLiteTokenRepository(path)
	s.Require().NoError(err)
	defer reopened.Close()

	found, err := reopened.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Equal(created, found)
}

func testIssuedToken() appidentity.IssuedToken {
	return appidentity.IssuedToken{
		ID:        "token-1",
		ExpiresAt: "2026-07-04T00:00:00Z",
		IssuedAt:  "2026-07-03T00:00:00Z",
		Methods:   []string{"password"},
		User: tokens.User{
			ID:   "user-1",
			Name: "admin",
			Domain: tokens.Domain{
				ID:   "default",
				Name: "Default",
			},
		},
		Project: tokens.Project{
			ID:   "demo",
			Name: "demo",
			Domain: tokens.Domain{
				ID:   "default",
				Name: "Default",
			},
		},
		Roles: []tokens.Role{
			{
				ID:   "admin",
				Name: "admin",
			},
		},
		Catalog: []tokens.CatalogEntry{},
	}
}
