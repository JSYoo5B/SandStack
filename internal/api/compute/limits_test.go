package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/limits"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
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
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *LimitsSuite) TearDownTest() {
	s.server.Close()
}

func (s *LimitsSuite) TestGetLimits() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	created, err := servers.Create(
		s.T().Context(),
		client,
		servers.CreateOpts{
			Name:      "test-server",
			ImageRef:  "img-1",
			FlavorRef: "1",
		},
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	result, err := limits.Get(
		s.T().Context(),
		client,
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(result)

	s.Assert().Equal(100, result.Absolute.MaxTotalInstances)
	s.Assert().Equal(1, result.Absolute.TotalInstancesUsed)
	s.Assert().Equal(1, result.Absolute.TotalCoresUsed)
	s.Assert().Equal(2048, result.Absolute.TotalRAMUsed)
}
