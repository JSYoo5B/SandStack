package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	"github.com/stretchr/testify/suite"
)

type NetworkSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestNetworkSuite(t *testing.T) {
	suite.Run(t, new(NetworkSuite))
}

func (s *NetworkSuite) SetupTest() {
	s.server = httptest.NewServer(
		network.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *NetworkSuite) TearDownTest() {
	s.server.Close()
}

func (s *NetworkSuite) TestListNetworks() {
	list := s.listNetworks()

	s.Assert().Empty(list)
}

func (s *NetworkSuite) TestCreateNetworkThenListNetworks() {
	created, err := networks.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		networks.CreateOpts{
			Name:      "private",
			ProjectID: "demo",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	list := s.listNetworks()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal("private", created.Name)
	s.Assert().Equal("ACTIVE", created.Status)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
	s.Assert().Equal("private", list[0].Name)
}

func (s *NetworkSuite) listNetworks() []networks.Network {
	pages, err := networks.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := networks.ExtractNetworks(pages)
	s.Require().NoError(err)

	return list
}
