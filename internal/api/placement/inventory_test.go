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

type InventorySuite struct {
	suite.Suite
	server *httptest.Server
}

func TestInventorySuite(t *testing.T) {
	suite.Run(t, new(InventorySuite))
}

func (s *InventorySuite) SetupTest() {
	s.server = httptest.NewServer(
		placement.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *InventorySuite) TearDownTest() {
	s.server.Close()
}

func (s *InventorySuite) TestUpdateAndGetInventories() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)

	updated, err := resourceproviders.UpdateInventories(
		context.Background(),
		client,
		provider.UUID,
		resourceproviders.UpdateInventoriesOpts{
			ResourceProviderGeneration: provider.Generation,
			Inventories: map[string]resourceproviders.Inventory{
				"VCPU": {
					AllocationRatio: 16.0,
					MaxUnit:         4,
					MinUnit:         1,
					Reserved:        0,
					StepSize:        1,
					Total:           4,
				},
			},
		},
	).Extract()
	s.Require().NoError(err)

	found, err := resourceproviders.GetInventories(
		context.Background(),
		client,
		provider.UUID,
	).Extract()
	s.Require().NoError(err)

	s.Require().Contains(updated.Inventories, "VCPU")
	s.Assert().Equal(4, updated.Inventories["VCPU"].Total)
	s.Require().Contains(found.Inventories, "VCPU")
	s.Assert().Equal(float32(16.0), found.Inventories["VCPU"].AllocationRatio)
}

func (s *InventorySuite) TestUpdateAndGetInventory() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)

	updated, err := resourceproviders.UpdateInventory(
		context.Background(),
		client,
		provider.UUID,
		"MEMORY_MB",
		resourceproviders.UpdateInventoryOpts{
			ResourceProviderGeneration: provider.Generation,
			Inventory: resourceproviders.Inventory{
				AllocationRatio: 1.5,
				MaxUnit:         2048,
				MinUnit:         1,
				Reserved:        128,
				StepSize:        1,
				Total:           2048,
			},
		},
	).Extract()
	s.Require().NoError(err)

	found, err := resourceproviders.GetInventory(
		context.Background(),
		client,
		provider.UUID,
		"MEMORY_MB",
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal(2048, updated.Total)
	s.Assert().Equal(128, found.Reserved)
	s.Assert().Equal(float32(1.5), found.AllocationRatio)
}

func (s *InventorySuite) TestDeleteInventory() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)

	_, err := resourceproviders.UpdateInventory(
		context.Background(),
		client,
		provider.UUID,
		"VCPU",
		resourceproviders.UpdateInventoryOpts{
			ResourceProviderGeneration: provider.Generation,
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

	err = resourceproviders.DeleteInventory(
		context.Background(),
		client,
		provider.UUID,
		"VCPU",
	).ExtractErr()
	s.Require().NoError(err)

	_, err = resourceproviders.GetInventory(
		context.Background(),
		client,
		provider.UUID,
		"VCPU",
	).Extract()
	s.Assert().Error(err)
}

func (s *InventorySuite) TestDeleteInventories() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)

	_, err := resourceproviders.UpdateInventory(
		context.Background(),
		client,
		provider.UUID,
		"VCPU",
		resourceproviders.UpdateInventoryOpts{
			ResourceProviderGeneration: provider.Generation,
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

	err = resourceproviders.DeleteInventories(
		context.Background(),
		client,
		provider.UUID,
	).ExtractErr()
	s.Require().NoError(err)

	found, err := resourceproviders.GetInventories(
		context.Background(),
		client,
		provider.UUID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().Empty(found.Inventories)
}

func (s *InventorySuite) createResourceProvider(
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
