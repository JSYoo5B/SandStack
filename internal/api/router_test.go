package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
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
