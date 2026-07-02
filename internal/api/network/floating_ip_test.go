package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/stretchr/testify/suite"
)

type FloatingIPSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestFloatingIPSuite(t *testing.T) {
	suite.Run(t, new(FloatingIPSuite))
}

func (s *FloatingIPSuite) SetupTest() {
	s.server = httptest.NewServer(
		network.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *FloatingIPSuite) TearDownTest() {
	s.server.Close()
}

func (s *FloatingIPSuite) TestListFloatingIPs() {
	list := s.listFloatingIPs()

	s.Assert().Empty(list)
}

func (s *FloatingIPSuite) TestCreateFloatingIPThenListFloatingIPs() {
	created := s.createFloatingIP("public")

	list := s.listFloatingIPs()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal("public", created.FloatingNetworkID)
	s.Assert().Equal("203.0.113.20", created.FloatingIP)
	s.Assert().Equal("DOWN", created.Status)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
}

func (s *FloatingIPSuite) TestGetFloatingIP() {
	created := s.createFloatingIP("public")

	found, err := floatingips.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.ID, found.ID)
	s.Assert().Equal("public", found.FloatingNetworkID)
}

func (s *FloatingIPSuite) TestDeleteFloatingIP() {
	created := s.createFloatingIP("public")

	err := floatingips.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listFloatingIPs()

	s.Assert().Empty(list)
}

func (s *FloatingIPSuite) listFloatingIPs() []floatingips.FloatingIP {
	pages, err := floatingips.List(
		testhelper.ServiceClient(s.server.URL),
		floatingips.ListOpts{},
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := floatingips.ExtractFloatingIPs(pages)
	s.Require().NoError(err)

	return list
}

func (s *FloatingIPSuite) createFloatingIP(
	floatingNetworkID string,
) *floatingips.FloatingIP {
	created, err := floatingips.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		floatingips.CreateOpts{
			Description:       "floating IP for tests",
			FloatingNetworkID: floatingNetworkID,
			FloatingIP:        "203.0.113.20",
			ProjectID:         "demo",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
