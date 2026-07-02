package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storecompute "github.com/JSYoo5B/SandStack/internal/store/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/stretchr/testify/suite"
)

type ActionSuite struct {
	suite.Suite
	server  *httptest.Server
	service *appcompute.Service
}

func TestActionSuite(t *testing.T) {
	suite.Run(t, new(ActionSuite))
}

func (s *ActionSuite) SetupTest() {
	s.service = appcompute.NewServiceWithRuntime(
		storecompute.NewMemoryServerRepository(),
		clock.Wall(),
		idgen.Random(),
	)
	s.server = httptest.NewServer(
		compute.NewRouterWithService(
			testhelper.DefaultConfig(),
			s.service,
		),
	)
}

func (s *ActionSuite) TearDownTest() {
	s.server.Close()
}

func (s *ActionSuite) TestStopAndStartServer() {
	created := s.service.CreateServer(appcompute.CreateServer{
		Name:     "web",
		ImageID:  "img-1",
		FlavorID: "1",
	})
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	err := servers.Stop(
		s.T().Context(),
		client,
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	stopped, err := s.service.GetServer(created.ID)
	s.Require().NoError(err)
	s.Assert().Equal("SHUTOFF", stopped.Status)

	err = servers.Start(
		s.T().Context(),
		client,
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	started, err := s.service.GetServer(created.ID)
	s.Require().NoError(err)
	s.Assert().Equal("ACTIVE", started.Status)
}

func (s *ActionSuite) TestRebootServer() {
	created := s.service.CreateServer(appcompute.CreateServer{
		Name:     "web",
		ImageID:  "img-1",
		FlavorID: "1",
	})
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	err := servers.Reboot(
		s.T().Context(),
		client,
		created.ID,
		servers.RebootOpts{Type: servers.SoftReboot},
	).ExtractErr()
	s.Require().NoError(err)

	rebooted, err := s.service.GetServer(created.ID)
	s.Require().NoError(err)

	s.Assert().Equal("ACTIVE", rebooted.Status)
	s.Assert().Equal(100, rebooted.Progress)
}
