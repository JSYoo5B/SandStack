package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/subnets"
	"github.com/stretchr/testify/suite"
)

type RouterInterfaceSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestRouterInterfaceSuite(t *testing.T) {
	suite.Run(t, new(RouterInterfaceSuite))
}

func (s *RouterInterfaceSuite) SetupTest() {
	s.server = httptest.NewServer(
		network.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *RouterInterfaceSuite) TearDownTest() {
	s.server.Close()
}

func (s *RouterInterfaceSuite) TestAddAndRemoveRouterInterface() {
	network := s.createNetwork("private")
	subnet := s.createSubnet(network.ID, "private-subnet")
	router := s.createRouter("edge")

	added, err := routers.AddInterface(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		router.ID,
		routers.AddInterfaceOpts{SubnetID: subnet.ID},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(added)

	removed, err := routers.RemoveInterface(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		router.ID,
		routers.RemoveInterfaceOpts{SubnetID: subnet.ID},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(removed)

	s.Assert().Equal(subnet.ID, added.SubnetID)
	s.Assert().NotEmpty(added.PortID)
	s.Assert().NotEmpty(added.ID)
	s.Assert().Equal(added.ID, removed.ID)
	s.Assert().Equal(added.PortID, removed.PortID)
	s.Assert().Equal("demo", removed.TenantID)
}

func (s *RouterInterfaceSuite) createNetwork(name string) *networks.Network {
	created, err := networks.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		networks.CreateOpts{
			Name:      name,
			ProjectID: "demo",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *RouterInterfaceSuite) createSubnet(
	networkID string,
	name string,
) *subnets.Subnet {
	created, err := subnets.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		subnets.CreateOpts{
			NetworkID:  networkID,
			Name:       name,
			CIDR:       "192.168.10.0/24",
			IPVersion:  gophercloud.IPv4,
			ProjectID:  "demo",
			EnableDHCP: boolPtr(true),
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *RouterInterfaceSuite) createRouter(name string) *routers.Router {
	created, err := routers.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		routers.CreateOpts{
			Name:      name,
			ProjectID: "demo",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
