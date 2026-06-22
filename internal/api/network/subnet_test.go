package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
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
	list := s.listSubnets()

	s.Assert().Empty(list)
}

func (s *SubnetSuite) TestCreateSubnetThenListSubnets() {
	network := s.createNetwork("private")

	created := s.createSubnet(network.ID, "private-subnet")
	list := s.listSubnets()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal(network.ID, created.NetworkID)
	s.Assert().Equal("private-subnet", created.Name)
	s.Assert().Equal("192.168.10.0/24", created.CIDR)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
	s.Assert().Equal(network.ID, list[0].NetworkID)
}

func (s *SubnetSuite) TestGetSubnet() {
	network := s.createNetwork("private")
	created := s.createSubnet(network.ID, "private-subnet")

	found, err := subnets.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.ID, found.ID)
	s.Assert().Equal(network.ID, found.NetworkID)
	s.Assert().Equal("private-subnet", found.Name)
}

func (s *SubnetSuite) TestDeleteSubnet() {
	network := s.createNetwork("private")
	created := s.createSubnet(network.ID, "private-subnet")

	err := subnets.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listSubnets()

	s.Assert().Empty(list)
}

func (s *SubnetSuite) listSubnets() []subnets.Subnet {
	pages, err := subnets.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := subnets.ExtractSubnets(pages)
	s.Require().NoError(err)

	return list
}

func (s *SubnetSuite) createNetwork(name string) *networks.Network {
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

func (s *SubnetSuite) createSubnet(
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

func boolPtr(value bool) *bool {
	return &value
}
