package volume_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/backups"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/stretchr/testify/suite"
)

type BackupSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestBackupSuite(t *testing.T) {
	suite.Run(t, new(BackupSuite))
}

func (s *BackupSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *BackupSuite) TearDownTest() {
	s.server.Close()
}

func (s *BackupSuite) TestCreateBackupThenListBackups() {
	volume := s.createVolume("database")

	created := s.createBackup(volume.ID, "database-backup")
	list := s.listBackups()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal(volume.ID, created.VolumeID)
	s.Assert().Equal("database-backup", created.Name)
	s.Assert().Equal("available", created.Status)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
}

func (s *BackupSuite) TestGetBackup() {
	volume := s.createVolume("database")
	created := s.createBackup(volume.ID, "database-backup")

	found, err := backups.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.ID, found.ID)
	s.Assert().Equal(volume.ID, found.VolumeID)
}

func (s *BackupSuite) TestDeleteBackup() {
	volume := s.createVolume("database")
	created := s.createBackup(volume.ID, "database-backup")

	err := backups.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listBackups()

	s.Assert().Empty(list)
}

func (s *BackupSuite) listBackups() []backups.Backup {
	pages, err := backups.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := backups.ExtractBackups(pages)
	s.Require().NoError(err)

	return list
}

func (s *BackupSuite) createBackup(
	volumeID string,
	name string,
) *backups.Backup {
	created, err := backups.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		backups.CreateOpts{
			VolumeID: volumeID,
			Name:     name,
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *BackupSuite) createVolume(name string) *volumes.Volume {
	created, err := volumes.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		volumes.CreateOpts{
			Size: 1,
			Name: name,
		},
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
