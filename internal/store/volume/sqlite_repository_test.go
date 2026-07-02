package volume_test

import (
	"path/filepath"
	"testing"

	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	storevolume "github.com/JSYoo5B/SandStack/internal/store/volume"
	"github.com/stretchr/testify/suite"
)

type SQLiteRepositorySuite struct {
	suite.Suite
	repository *storevolume.SQLiteRepository
}

func TestSQLiteRepositorySuite(t *testing.T) {
	suite.Run(t, new(SQLiteRepositorySuite))
}

func (s *SQLiteRepositorySuite) SetupTest() {
	repository, err := storevolume.OpenSQLiteRepository(":memory:")
	s.Require().NoError(err)

	s.repository = repository
}

func (s *SQLiteRepositorySuite) TearDownTest() {
	s.Require().NoError(s.repository.Close())
}

func (s *SQLiteRepositorySuite) TestCreateListGetAndUpdateVolume() {
	created := s.repository.Create(volumeFixture("vol-1"))
	created.Status = "available"

	updated, err := s.repository.Update(created)
	s.Require().NoError(err)
	listed := s.repository.List()
	found, err := s.repository.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Len(listed, 1)
	s.Assert().Equal(updated, listed[0])
	s.Assert().Equal(updated, found)
}

func (s *SQLiteRepositorySuite) TestDeleteVolume() {
	created := s.repository.Create(volumeFixture("vol-1"))

	err := s.repository.Delete(created.ID)
	s.Require().NoError(err)

	_, err = s.repository.Get(created.ID)
	s.Require().ErrorIs(err, appvolume.ErrVolumeNotFound)
	s.Assert().Empty(s.repository.List())
}

func (s *SQLiteRepositorySuite) TestResetClearsVolumes() {
	s.repository.Create(volumeFixture("vol-1"))

	s.repository.Reset()

	s.Assert().Empty(s.repository.List())
}

func (s *SQLiteRepositorySuite) TestFileBackedDatabasePersistsVolumes() {
	path := filepath.Join(s.T().TempDir(), "sandstack.db")
	repository, err := storevolume.OpenSQLiteRepository(path)
	s.Require().NoError(err)

	created := repository.Create(volumeFixture("vol-1"))
	s.Require().NoError(repository.Close())

	reopened, err := storevolume.OpenSQLiteRepository(path)
	s.Require().NoError(err)
	defer reopened.Close()

	found, err := reopened.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Equal(created, found)
}

func volumeFixture(id string) appvolume.Volume {
	return appvolume.Volume{
		ID:          id,
		Status:      "creating",
		Size:        1,
		Name:        "database",
		Description: "test volume",
		VolumeType:  "__DEFAULT__",
		Metadata: map[string]string{
			"role": "database",
		},
		CreatedAt:   "2026-07-03T00:00:00.000000",
		UpdatedAt:   "2026-07-03T00:00:00.000000",
		Bootable:    "false",
		Encrypted:   true,
		Multiattach: true,
	}
}
