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

type AggregateSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestAggregateSuite(t *testing.T) {
	suite.Run(t, new(AggregateSuite))
}

func (s *AggregateSuite) SetupTest() {
	s.server = httptest.NewServer(
		placement.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *AggregateSuite) TearDownTest() {
	s.server.Close()
}

func (s *AggregateSuite) TestUpdateAndGetAggregates() {
	client := testhelper.ServiceClient(s.server.URL)
	provider := s.createResourceProvider(client)

	updated, err := resourceproviders.UpdateAggregates(
		context.Background(),
		client,
		provider.UUID,
		resourceproviders.UpdateAggregatesOpts{
			ResourceProviderGeneration: &provider.Generation,
			Aggregates: []string{
				"aggregate-1",
				"aggregate-2",
			},
		},
	).Extract()
	s.Require().NoError(err)

	found, err := resourceproviders.GetAggregates(
		context.Background(),
		client,
		provider.UUID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().ElementsMatch(
		[]string{"aggregate-1", "aggregate-2"},
		updated.Aggregates,
	)
	s.Assert().ElementsMatch(updated.Aggregates, found.Aggregates)
}

func (s *AggregateSuite) createResourceProvider(
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
