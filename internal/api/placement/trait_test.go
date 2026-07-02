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

type TraitSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestTraitSuite(t *testing.T) {
	suite.Run(t, new(TraitSuite))
}

func (s *TraitSuite) SetupTest() {
	s.server = httptest.NewServer(
		placement.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *TraitSuite) TearDownTest() {
	s.server.Close()
}

func (s *TraitSuite) TestUpdateAndGetTraits() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)

	updated, err := resourceproviders.UpdateTraits(
		context.Background(),
		client,
		provider.UUID,
		resourceproviders.UpdateTraitsOpts{
			ResourceProviderGeneration: provider.Generation,
			Traits: []string{
				"CUSTOM_FAST_DISK",
				"CUSTOM_GPU",
			},
		},
	).Extract()
	s.Require().NoError(err)

	found, err := resourceproviders.GetTraits(
		context.Background(),
		client,
		provider.UUID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().ElementsMatch(
		[]string{"CUSTOM_FAST_DISK", "CUSTOM_GPU"},
		updated.Traits,
	)
	s.Assert().ElementsMatch(updated.Traits, found.Traits)
}

func (s *TraitSuite) TestDeleteTraits() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)

	_, err := resourceproviders.UpdateTraits(
		context.Background(),
		client,
		provider.UUID,
		resourceproviders.UpdateTraitsOpts{
			ResourceProviderGeneration: provider.Generation,
			Traits: []string{
				"CUSTOM_GPU",
			},
		},
	).Extract()
	s.Require().NoError(err)

	err = resourceproviders.DeleteTraits(
		context.Background(),
		client,
		provider.UUID,
	).ExtractErr()
	s.Require().NoError(err)

	found, err := resourceproviders.GetTraits(
		context.Background(),
		client,
		provider.UUID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().Empty(found.Traits)
}

func (s *TraitSuite) createResourceProvider(
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
