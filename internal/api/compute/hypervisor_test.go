package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/hypervisors"
	"github.com/stretchr/testify/suite"
)

type HypervisorSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestHypervisorSuite(t *testing.T) {
	suite.Run(t, new(HypervisorSuite))
}

func (s *HypervisorSuite) SetupTest() {
	s.server = httptest.NewServer(
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *HypervisorSuite) TearDownTest() {
	s.server.Close()
}

func (s *HypervisorSuite) TestListHypervisors() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	pages, err := hypervisors.List(
		client,
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	listed, err := hypervisors.ExtractHypervisors(pages)
	s.Require().NoError(err)

	s.Require().Len(listed, 1)
	s.Assert().Equal("1", listed[0].ID)
	s.Assert().Equal("sandstack", listed[0].HypervisorHostname)
	s.Assert().Equal("enabled", listed[0].Status)
}

func (s *HypervisorSuite) TestGetHypervisor() {
	found, err := hypervisors.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		"1",
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal("1", found.ID)
	s.Assert().Equal("sandstack", found.HypervisorHostname)
}

func (s *HypervisorSuite) TestGetHypervisorStatistics() {
	stats, err := hypervisors.GetStatistics(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(stats)

	s.Assert().Equal(1, stats.Count)
	s.Assert().Equal(200, stats.VCPUs)
	s.Assert().Equal(512000, stats.MemoryMB)
}
