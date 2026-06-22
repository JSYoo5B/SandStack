package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/stretchr/testify/suite"
)

type FaultCompatibilitySuite struct {
	suite.Suite
	server *httptest.Server
}

func TestFaultCompatibilitySuite(t *testing.T) {
	suite.Run(t, new(FaultCompatibilitySuite))
}

func (s *FaultCompatibilitySuite) SetupTest() {
	s.server = httptest.NewServer(api.NewRouter(testhelper.DefaultConfig()))
}

func (s *FaultCompatibilitySuite) TearDownTest() {
	s.server.Close()
}

func (s *FaultCompatibilitySuite) TestThirdServerCreateReturns503() {
	response, err := http.Post(
		s.server.URL+"/_sandstack/faults",
		"application/json",
		bytes.NewBufferString(serverCreateFaultBody()),
	)
	s.Require().NoError(err)
	defer response.Body.Close()
	s.Require().Equal(http.StatusCreated, response.StatusCode)

	s.createServer("web-1")
	s.createServer("web-2")

	_, err = servers.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/compute/v2.1/demo"),
		servers.CreateOpts{
			Name:      "web-3",
			ImageRef:  "img-1",
			FlavorRef: "1",
		},
		nil,
	).Extract()
	s.Require().Error(err)
	s.Assert().True(strings.Contains(err.Error(), "503"))
}

func (s *FaultCompatibilitySuite) createServer(name string) {
	created, err := servers.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/compute/v2.1/demo"),
		servers.CreateOpts{
			Name:      name,
			ImageRef:  "img-1",
			FlavorRef: "1",
		},
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)
}

func serverCreateFaultBody() string {
	return `{
		"rule": {
			"id": "third-server-create",
			"service": "compute",
			"operation": "server.create",
			"behavior": {
				"http_status": 503,
				"message": "injected server create failure"
			},
			"trigger": {
				"nth": 3,
				"once": true
			}
		}
	}`
}
