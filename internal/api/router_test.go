package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack"
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
	s.server = httptest.NewServer(api.NewRouter(config.Load()))
}

func (s *RouterSuite) TearDownTest() {
	s.server.Close()
}

func (s *RouterSuite) TestRootIsNotExposed() {
	response, err := http.Get(s.server.URL + "/")
	s.Require().NoError(err)
	defer response.Body.Close()

	s.Assert().Equal(http.StatusNotFound, response.StatusCode)
}

func (s *RouterSuite) TestMountedIdentityPasswordAuth() {
	authOptions := testhelper.PasswordAuthOptions()
	authOptions.IdentityEndpoint = s.server.URL + "/identity/v3"

	provider, err := openstack.AuthenticatedClient(
		s.T().Context(),
		*authOptions,
	)
	s.Require().NoError(err)

	s.Assert().NotEmpty(provider.TokenID)
}
