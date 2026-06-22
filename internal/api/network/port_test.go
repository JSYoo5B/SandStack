package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
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
	pages, err := ports.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := ports.ExtractPorts(pages)
	s.Require().NoError(err)

	s.Assert().Empty(list)
}
