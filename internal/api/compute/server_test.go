package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/stretchr/testify/suite"
)

type ServerSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

func (s *ServerSuite) SetupTest() {
	s.server = httptest.NewServer(
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ServerSuite) TearDownTest() {
	s.server.Close()
}

func (s *ServerSuite) TestListServers() {
	pages, err := servers.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := servers.ExtractServers(pages)
	s.Require().NoError(err)

	s.Assert().Empty(list)
}

func (s *ServerSuite) TestListSimpleServers() {
	pages, err := servers.ListSimple(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := servers.ExtractServers(pages)
	s.Require().NoError(err)

	s.Assert().Empty(list)
}

func (s *ServerSuite) TestCreateServerThenListServers() {
	created := s.createServer("test-server")

	pages, err := servers.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := servers.ExtractServers(pages)
	s.Require().NoError(err)

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal("test-server", created.Name)
	s.Assert().Equal("BUILD", created.Status)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
	s.Assert().Equal("test-server", list[0].Name)
}

func (s *ServerSuite) createServer(name string) *servers.Server {
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
