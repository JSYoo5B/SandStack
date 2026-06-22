package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/subnets"
	"github.com/stretchr/testify/suite"
)

type SubnetSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestSubnetSuite(t *testing.T) {
	suite.Run(t, new(SubnetSuite))
}

func (s *SubnetSuite) SetupTest() {
	s.server = httptest.NewServer(
		network.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *SubnetSuite) TearDownTest() {
	s.server.Close()
}

func (s *SubnetSuite) TestListSubnets() {
	pages, err := subnets.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := subnets.ExtractSubnets(pages)
	s.Require().NoError(err)

	s.Assert().Empty(list)
}
