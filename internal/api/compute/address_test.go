package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/stretchr/testify/suite"
)

type AddressSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestAddressSuite(t *testing.T) {
	suite.Run(t, new(AddressSuite))
}

func (s *AddressSuite) SetupTest() {
	s.server = httptest.NewServer(
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *AddressSuite) TearDownTest() {
	s.server.Close()
}

func (s *AddressSuite) TestListServerAddresses() {
	created := s.createServer("web")

	pages, err := servers.ListAddresses(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	addresses, err := servers.ExtractAddresses(pages)
	s.Require().NoError(err)
	s.Require().Contains(addresses, "private")
	s.Require().Len(addresses["private"], 1)
	s.Assert().Equal(4, addresses["private"][0].Version)
	s.Assert().Equal("10.0.0.10", addresses["private"][0].Address)
}

func (s *AddressSuite) TestListServerAddressesByNetwork() {
	created := s.createServer("web")

	pages, err := servers.ListAddressesByNetwork(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
		"private",
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	addresses, err := servers.ExtractNetworkAddresses(pages)
	s.Require().NoError(err)
	s.Require().Len(addresses, 1)
	s.Assert().Equal(4, addresses[0].Version)
	s.Assert().Equal("10.0.0.10", addresses[0].Address)
}

func (s *AddressSuite) createServer(name string) *servers.Server {
	created, err := servers.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		servers.CreateOpts{
			Name:      name,
			ImageRef:  "img-1",
			FlavorRef: "1",
		},
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
