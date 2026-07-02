package placement_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/placement"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/placement/v1/resourceproviders"
	"github.com/stretchr/testify/suite"
)

type UsageSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestUsageSuite(t *testing.T) {
	suite.Run(t, new(UsageSuite))
}

func (s *UsageSuite) SetupTest() {
	s.server = httptest.NewServer(
		placement.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *UsageSuite) TearDownTest() {
	s.server.Close()
}

func (s *UsageSuite) TestGetUsages() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)

	usages, err := resourceproviders.GetUsages(
		context.Background(),
		client,
		provider.UUID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal(provider.Generation, usages.ResourceProviderGeneration)
	s.Assert().Empty(usages.Usages)
}

func (s *UsageSuite) createResourceProvider(
	client *gophercloud.ServiceClient,
) *resourceproviders.ResourceProvider {
	provider, err := resourceproviders.Create(
		context.Background(),
		client,
		resourceproviders.CreateOpts{
			Name: "compute-1",
			UUID: "resource-provider-1",
		},
	).Extract()
	s.Require().NoError(err)

	return provider
}
