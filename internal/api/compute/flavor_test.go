package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/flavors"
	"github.com/stretchr/testify/suite"
)

type FlavorSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestFlavorSuite(t *testing.T) {
	suite.Run(t, new(FlavorSuite))
}

func (s *FlavorSuite) SetupTest() {
	s.server = httptest.NewServer(
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *FlavorSuite) TearDownTest() {
	s.server.Close()
}

func (s *FlavorSuite) TestListFlavors() {
	pages, err := flavors.ListDetail(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := flavors.ExtractFlavors(pages)
	s.Require().NoError(err)

	s.Require().Len(list, 1)
	s.Assert().Equal("1", list[0].ID)
	s.Assert().Equal("m1.small", list[0].Name)
	s.Assert().Equal(2048, list[0].RAM)
	s.Assert().Equal(1, list[0].VCPUs)
	s.Assert().Equal(20, list[0].Disk)
}

func (s *FlavorSuite) TestGetFlavor() {
	flavor, err := flavors.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		"1",
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(flavor)

	s.Assert().Equal("1", flavor.ID)
	s.Assert().Equal("m1.small", flavor.Name)
	s.Assert().Equal(true, flavor.IsPublic)
}
