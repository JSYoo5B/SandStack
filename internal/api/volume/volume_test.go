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
	pages, err := volumes.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := volumes.ExtractVolumes(pages)
	s.Require().NoError(err)

	s.Assert().Empty(list)
}
