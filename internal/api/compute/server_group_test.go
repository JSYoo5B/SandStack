package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servergroups"
	"github.com/stretchr/testify/suite"
)

type ServerGroupSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestServerGroupSuite(t *testing.T) {
	suite.Run(t, new(ServerGroupSuite))
}

func (s *ServerGroupSuite) SetupTest() {
	s.server = httptest.NewServer(
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ServerGroupSuite) TearDownTest() {
	s.server.Close()
}

func (s *ServerGroupSuite) TestCreateListGetAndDeleteServerGroup() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	created, err := servergroups.Create(
		s.T().Context(),
		client,
		servergroups.CreateOpts{
			Name:     "web-group",
			Policies: []string{"anti-affinity"},
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	pages, err := servergroups.List(
		client,
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)
	listed, err := servergroups.ExtractServerGroups(pages)
	s.Require().NoError(err)
	found, err := servergroups.Get(
		s.T().Context(),
		client,
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal("web-group", created.Name)
	s.Assert().Equal([]string{"anti-affinity"}, created.Policies)
	s.Require().Len(listed, 1)
	s.Assert().Equal(created.ID, listed[0].ID)
	s.Assert().Equal(created.ID, found.ID)

	err = servergroups.Delete(
		s.T().Context(),
		client,
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)
}
