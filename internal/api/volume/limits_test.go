package volume_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/limits"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/stretchr/testify/suite"
)

type LimitsSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestLimitsSuite(t *testing.T) {
	suite.Run(t, new(LimitsSuite))
}

func (s *LimitsSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *LimitsSuite) TearDownTest() {
	s.server.Close()
}

func (s *LimitsSuite) TestGetLimits() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	created, err := volumes.Create(
		s.T().Context(),
		client,
		volumes.CreateOpts{
			Size: 7,
			Name: "database",
		},
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	result, err := limits.Get(s.T().Context(), client).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(result)

	s.Assert().Equal(1000, result.Absolute.MaxTotalVolumes)
	s.Assert().Equal(1, result.Absolute.TotalVolumesUsed)
	s.Assert().Equal(7, result.Absolute.TotalGigabytesUsed)
	s.Assert().Empty(result.Rate)
}
