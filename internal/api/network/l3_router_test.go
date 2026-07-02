package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/routers"
	"github.com/stretchr/testify/suite"
)

type RouterSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterSuite))
}

func (s *RouterSuite) SetupTest() {
	s.server = httptest.NewServer(
		network.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *RouterSuite) TearDownTest() {
	s.server.Close()
}

func (s *RouterSuite) TestListRouters() {
	list := s.listRouters()

	s.Assert().Empty(list)
}

func (s *RouterSuite) TestCreateRouterThenListRouters() {
	created := s.createRouter("edge")

	list := s.listRouters()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal("edge", created.Name)
	s.Assert().Equal("ACTIVE", created.Status)
	s.Assert().True(created.AdminStateUp)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
	s.Assert().Equal("edge", list[0].Name)
}

func (s *RouterSuite) TestGetRouter() {
	created := s.createRouter("edge")

	found, err := routers.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.ID, found.ID)
	s.Assert().Equal("edge", found.Name)
}

func (s *RouterSuite) TestDeleteRouter() {
	created := s.createRouter("edge")

	err := routers.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listRouters()

	s.Assert().Empty(list)
}

func (s *RouterSuite) listRouters() []routers.Router {
	pages, err := routers.List(
		testhelper.ServiceClient(s.server.URL),
		routers.ListOpts{},
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := routers.ExtractRouters(pages)
	s.Require().NoError(err)

	return list
}

func (s *RouterSuite) createRouter(name string) *routers.Router {
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
