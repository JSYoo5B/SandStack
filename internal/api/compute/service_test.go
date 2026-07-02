package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/services"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	s.server = httptest.NewServer(
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ServiceSuite) TearDownTest() {
	s.server.Close()
}

func (s *ServiceSuite) TestListComputeServices() {
	pages, err := services.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	listed, err := services.ExtractServices(pages)
	s.Require().NoError(err)

	s.Require().Len(listed, 2)
	s.Assert().Equal("compute-service-1", listed[0].ID)
	s.Assert().Equal("nova-compute", listed[0].Binary)
	s.Assert().Equal("up", listed[0].State)
	s.Assert().Equal("enabled", listed[0].Status)
}
