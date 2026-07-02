package identity_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/endpoints"
	"github.com/stretchr/testify/suite"
)

type EndpointSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestEndpointSuite(t *testing.T) {
	suite.Run(t, new(EndpointSuite))
}

func (s *EndpointSuite) SetupTest() {
	s.server = httptest.NewServer(
		identity.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *EndpointSuite) TearDownTest() {
	s.server.Close()
}

func (s *EndpointSuite) TestListEndpoints() {
	pages, err := endpoints.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	listed, err := endpoints.ExtractEndpoints(pages)
	s.Require().NoError(err)

	s.Assert().Contains(
		identityEndpointIDs(listed),
		"compute-public",
	)
}

func (s *EndpointSuite) TestGetEndpoint() {
	result := endpoints.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		"compute-public",
	)
	endpoint, err := result.Extract()
	s.Require().NoError(err)
	s.Require().NotNil(endpoint)

	s.Assert().Equal("compute-public", endpoint.ID)
	s.Assert().Equal("compute", endpoint.ServiceID)
	s.Assert().Equal(gophercloud.AvailabilityPublic, endpoint.Availability)
	s.Assert().Equal(s.server.URL+"/compute/v2.1/demo", endpoint.URL)
}

func identityEndpointIDs(listed []endpoints.Endpoint) []string {
	ids := make([]string, 0, len(listed))
	for _, endpoint := range listed {
		ids = append(ids, endpoint.ID)
	}

	return ids
}
