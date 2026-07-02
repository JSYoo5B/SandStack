package volume_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/snapshots"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/stretchr/testify/suite"
)

type SnapshotSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestSnapshotSuite(t *testing.T) {
	suite.Run(t, new(SnapshotSuite))
}

func (s *SnapshotSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *SnapshotSuite) TearDownTest() {
	s.server.Close()
}

func (s *SnapshotSuite) TestListSnapshots() {
	list := s.listSnapshots()

	s.Assert().Empty(list)
}

func (s *SnapshotSuite) TestCreateSnapshotThenListSnapshots() {
	volume := s.createVolume("database")

	created := s.createSnapshot(volume.ID, "database-snapshot")
	list := s.listSnapshots()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal(volume.ID, created.VolumeID)
	s.Assert().Equal("database-snapshot", created.Name)
	s.Assert().Equal("available", created.Status)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
}

func (s *SnapshotSuite) TestGetSnapshot() {
	volume := s.createVolume("database")
	created := s.createSnapshot(volume.ID, "database-snapshot")

	found, err := snapshots.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.ID, found.ID)
	s.Assert().Equal(volume.ID, found.VolumeID)
}

func (s *SnapshotSuite) TestDeleteSnapshot() {
	volume := s.createVolume("database")
	created := s.createSnapshot(volume.ID, "database-snapshot")

	err := snapshots.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listSnapshots()

	s.Assert().Empty(list)
}

func (s *SnapshotSuite) listSnapshots() []snapshots.Snapshot {
	pages, err := snapshots.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := snapshots.ExtractSnapshots(pages)
	s.Require().NoError(err)

	return list
}

func (s *SnapshotSuite) createSnapshot(
	volumeID string,
	name string,
) *snapshots.Snapshot {
	created, err := snapshots.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		snapshots.CreateOpts{
			VolumeID: volumeID,
			Name:     name,
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *SnapshotSuite) createVolume(name string) *volumes.Volume {
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
