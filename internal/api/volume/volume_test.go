package volume_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/stretchr/testify/suite"
)

type VolumeSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestVolumeSuite(t *testing.T) {
	suite.Run(t, new(VolumeSuite))
}

func (s *VolumeSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *VolumeSuite) TearDownTest() {
	s.server.Close()
}

func (s *VolumeSuite) TestListVolumes() {
	list := s.listVolumes()

	s.Assert().Empty(list)
}

func (s *VolumeSuite) TestCreateVolumeThenListVolumes() {
	created := s.createVolume("database")

	list := s.listVolumes()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal("database", created.Name)
	s.Assert().Equal(1, created.Size)
	s.Assert().Equal("creating", created.Status)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
	s.Assert().Equal("database", list[0].Name)
}

func (s *VolumeSuite) TestGetVolume() {
	created := s.createVolume("database")

	found, err := volumes.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.ID, found.ID)
	s.Assert().Equal("database", found.Name)
}

func (s *VolumeSuite) TestDeleteVolume() {
	created := s.createVolume("database")

	err := volumes.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
		volumes.DeleteOpts{},
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listVolumes()

	s.Assert().Empty(list)
}

func (s *VolumeSuite) listVolumes() []volumes.Volume {
	pages, err := volumes.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := volumes.ExtractVolumes(pages)
	s.Require().NoError(err)

	return list
}

func (s *VolumeSuite) createVolume(name string) *volumes.Volume {
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
