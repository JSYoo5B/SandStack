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

type AllocationSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestAllocationSuite(t *testing.T) {
	suite.Run(t, new(AllocationSuite))
}

func (s *AllocationSuite) SetupTest() {
	s.server = httptest.NewServer(
		placement.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *AllocationSuite) TearDownTest() {
	s.server.Close()
}

func (s *AllocationSuite) TestGetAllocations() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)

	allocations, err := resourceproviders.GetAllocations(
		context.Background(),
		client,
		provider.UUID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal(
		provider.Generation,
		allocations.ResourceProviderGeneration,
	)
	s.Assert().Empty(allocations.Allocations)
}

func (s *AllocationSuite) createResourceProvider(
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
