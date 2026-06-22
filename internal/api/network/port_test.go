package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
	"github.com/stretchr/testify/suite"
)

type PortSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestPortSuite(t *testing.T) {
	suite.Run(t, new(PortSuite))
}

func (s *PortSuite) SetupTest() {
	s.server = httptest.NewServer(
		network.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *PortSuite) TearDownTest() {
	s.server.Close()
}

func (s *PortSuite) TestListPorts() {
	list := s.listPorts()

	s.Assert().Empty(list)
}

func (s *PortSuite) TestCreatePortThenListPorts() {
	network := s.createNetwork("private")

	created := s.createPort(network.ID, "private-port")
	list := s.listPorts()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal(network.ID, created.NetworkID)
	s.Assert().Equal("private-port", created.Name)
	s.Assert().Equal("DOWN", created.Status)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
	s.Assert().Equal(network.ID, list[0].NetworkID)
}

func (s *PortSuite) listPorts() []ports.Port {
	pages, err := ports.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := ports.ExtractPorts(pages)
	s.Require().NoError(err)

	return list
}

func (s *PortSuite) createNetwork(name string) *networks.Network {
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

func (s *PortSuite) createPort(networkID string, name string) *ports.Port {
	created, err := ports.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		ports.CreateOpts{
			NetworkID:    networkID,
			Name:         name,
			AdminStateUp: boolPtr(true),
			ProjectID:    "demo",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
