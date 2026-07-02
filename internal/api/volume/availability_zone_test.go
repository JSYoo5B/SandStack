package volume_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/availabilityzones"
	"github.com/stretchr/testify/suite"
)

type AvailabilityZoneSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestAvailabilityZoneSuite(t *testing.T) {
	suite.Run(t, new(AvailabilityZoneSuite))
}

func (s *AvailabilityZoneSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *AvailabilityZoneSuite) TearDownTest() {
	s.server.Close()
}

func (s *AvailabilityZoneSuite) TestListAvailabilityZones() {
	pages, err := availabilityzones.List(
		testhelper.ServiceClient(s.server.URL + "/demo"),
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := availabilityzones.ExtractAvailabilityZones(pages)
	s.Require().NoError(err)

	s.Require().Len(list, 1)
	s.Assert().Equal("nova", list[0].ZoneName)
	s.Assert().True(list[0].ZoneState.Available)
}
