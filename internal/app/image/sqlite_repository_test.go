package image_test

import (
	"path/filepath"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/app/image"
	"github.com/stretchr/testify/suite"
)

type SQLiteRepositorySuite struct {
	suite.Suite
	repository *image.SQLiteRepository
}

func TestSQLiteRepositorySuite(t *testing.T) {
	suite.Run(t, new(SQLiteRepositorySuite))
}

func (s *SQLiteRepositorySuite) SetupTest() {
	repository, err := image.OpenSQLiteRepository(":memory:")
	s.Require().NoError(err)

	s.repository = repository
}

func (s *SQLiteRepositorySuite) TearDownTest() {
	s.Require().NoError(s.repository.Close())
}

func (s *SQLiteRepositorySuite) TestCreateListAndGetImage() {
	created := s.repository.Create(image.Image{
		ID:              "img-1",
		Name:            "ubuntu",
		Status:          "queued",
		ContainerFormat: "bare",
		DiskFormat:      "qcow2",
		MinDisk:         1,
		MinRAM:          2,
		Protected:       true,
		Visibility:      "private",
		Tags:            []string{"linux", "test"},
		CreatedAt:       "2026-07-03T00:00:00Z",
		UpdatedAt:       "2026-07-03T00:00:00Z",
	})
	created.Name = "ubuntu-updated"

	updated, err := s.repository.Update(created)
	s.Require().NoError(err)
	listed := s.repository.List()
	found, err := s.repository.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Len(listed, 1)
	s.Assert().Equal(updated, listed[0])
	s.Assert().Equal(updated, found)
}

func (s *SQLiteRepositorySuite) TestDeleteImage() {
	created := s.repository.Create(image.Image{
		ID:              "img-1",
		Name:            "ubuntu",
		Status:          "queued",
		ContainerFormat: "bare",
		DiskFormat:      "qcow2",
		Visibility:      "private",
		Tags:            []string{},
		CreatedAt:       "2026-07-03T00:00:00Z",
		UpdatedAt:       "2026-07-03T00:00:00Z",
	})

	err := s.repository.Delete(created.ID)
	s.Require().NoError(err)

	_, err = s.repository.Get(created.ID)
	s.Require().ErrorIs(err, image.ErrImageNotFound)
	s.Assert().Empty(s.repository.List())
}

func (s *SQLiteRepositorySuite) TestResetClearsImages() {
	s.repository.Create(image.Image{
		ID:              "img-1",
		Name:            "ubuntu",
		Status:          "queued",
		ContainerFormat: "bare",
		DiskFormat:      "qcow2",
		Visibility:      "private",
		Tags:            []string{},
		CreatedAt:       "2026-07-03T00:00:00Z",
		UpdatedAt:       "2026-07-03T00:00:00Z",
	})

	s.repository.Reset()

	s.Assert().Empty(s.repository.List())
}

func (s *SQLiteRepositorySuite) TestFileBackedDatabasePersistsImages() {
	path := filepath.Join(s.T().TempDir(), "sandstack.db")
	repository, err := image.OpenSQLiteRepository(path)
	s.Require().NoError(err)

	created := repository.Create(image.Image{
		ID:              "img-1",
		Name:            "ubuntu",
		Status:          "queued",
		ContainerFormat: "bare",
		DiskFormat:      "qcow2",
		Visibility:      "private",
		Tags:            []string{},
		CreatedAt:       "2026-07-03T00:00:00Z",
		UpdatedAt:       "2026-07-03T00:00:00Z",
	})
	s.Require().NoError(repository.Close())

	reopened, err := image.OpenSQLiteRepository(path)
	s.Require().NoError(err)
	defer reopened.Close()

	found, err := reopened.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Equal(created, found)
}
