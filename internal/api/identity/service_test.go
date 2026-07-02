package identity_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/services"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	s.server = httptest.NewServer(
		identity.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ServiceSuite) TearDownTest() {
	s.server.Close()
}

func (s *ServiceSuite) TestListServices() {
	pages, err := services.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	listed, err := services.ExtractServices(pages)
	s.Require().NoError(err)

	s.Assert().Equal(
		map[string]bool{
			"identity":  true,
			"compute":   true,
			"network":   true,
			"image":     true,
			"volumev3":  true,
			"placement": true,
		},
		identityServiceTypes(listed),
	)
}

func (s *ServiceSuite) TestGetService() {
	result := services.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		"compute",
	)
	service, err := result.Extract()
	s.Require().NoError(err)
	s.Require().NotNil(service)

	s.Assert().Equal("compute", service.ID)
	s.Assert().Equal("compute", service.Type)
	s.Assert().True(service.Enabled)
}

func identityServiceTypes(listed []services.Service) map[string]bool {
	types := make(map[string]bool, len(listed))
	for _, service := range listed {
		types[service.Type] = true
	}

	return types
}
