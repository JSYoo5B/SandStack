package volume_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumetypes"
	"github.com/stretchr/testify/suite"
)

type VolumeTypeSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestVolumeTypeSuite(t *testing.T) {
	suite.Run(t, new(VolumeTypeSuite))
}

func (s *VolumeTypeSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *VolumeTypeSuite) TearDownTest() {
	s.server.Close()
}

func (s *VolumeTypeSuite) TestListVolumeTypes() {
	pages, err := volumetypes.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := volumetypes.ExtractVolumeTypes(pages)
	s.Require().NoError(err)

	s.Require().Len(list, 1)
	s.Assert().Equal("default", list[0].ID)
	s.Assert().Equal("__DEFAULT__", list[0].Name)
	s.Assert().Equal(true, list[0].IsPublic)
}

func (s *VolumeTypeSuite) TestGetVolumeType() {
	volumeType, err := volumetypes.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		"default",
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(volumeType)

	s.Assert().Equal("default", volumeType.ID)
	s.Assert().Equal("__DEFAULT__", volumeType.Name)
}
