package placement_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/placement"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/placement/v1/allocationcandidates"
	"github.com/gophercloud/gophercloud/v2/openstack/placement/v1/resourceproviders"
	"github.com/stretchr/testify/suite"
)

type AllocationCandidateSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestAllocationCandidateSuite(t *testing.T) {
	suite.Run(t, new(AllocationCandidateSuite))
}

func (s *AllocationCandidateSuite) SetupTest() {
	s.server = httptest.NewServer(
		placement.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *AllocationCandidateSuite) TearDownTest() {
	s.server.Close()
}

func (s *AllocationCandidateSuite) TestListAllocationCandidates() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)
	s.updateInventory(client, provider.UUID)

	pages, err := allocationcandidates.List(
		client,
		allocationcandidates.ListOpts{
			Resources: "VCPU:2",
		},
	).AllPages(context.Background())
	s.Require().NoError(err)

	candidates, err := allocationcandidates.ExtractAllocationCandidates(pages)
	s.Require().NoError(err)

	s.Require().Len(candidates.AllocationRequests, 1)
	allocation := candidates.AllocationRequests[0].Allocations[provider.UUID]
	s.Assert().Equal(2, allocation.Resources["VCPU"])
	s.Require().Contains(candidates.ProviderSummaries, provider.UUID)
	s.Assert().Equal(
		64,
		candidates.ProviderSummaries[provider.UUID].Resources["VCPU"].Capacity,
	)
}

func (s *AllocationCandidateSuite) createResourceProvider(
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

func (s *AllocationCandidateSuite) updateInventory(
	client *gophercloud.ServiceClient,
	providerUUID string,
) {
	_, err := resourceproviders.UpdateInventory(
		context.Background(),
		client,
		providerUUID,
		"VCPU",
		resourceproviders.UpdateInventoryOpts{
			Inventory: resourceproviders.Inventory{
				AllocationRatio: 16.0,
				MaxUnit:         4,
				MinUnit:         1,
				Reserved:        0,
				StepSize:        1,
				Total:           4,
			},
		},
	).Extract()
	s.Require().NoError(err)
}
