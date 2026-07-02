package placement_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/placement"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/placement/v1/resourceproviders"
	"github.com/stretchr/testify/suite"
)

type ResourceProviderSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestResourceProviderSuite(t *testing.T) {
	suite.Run(t, new(ResourceProviderSuite))
}

func (s *ResourceProviderSuite) SetupTest() {
	s.server = httptest.NewServer(
		placement.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ResourceProviderSuite) TearDownTest() {
	s.server.Close()
}

func (s *ResourceProviderSuite) TestListEmptyResourceProviders() {
	client := testhelper.ServiceClient(s.server.URL)

	pages, err := resourceproviders.List(client, nil).AllPages(context.Background())
	s.Require().NoError(err)

	providers, err := resourceproviders.ExtractResourceProviders(pages)
	s.Require().NoError(err)

	s.Assert().Empty(providers)
}

func (s *ResourceProviderSuite) TestCreateAndGetResourceProvider() {
	client := testhelper.ServiceClient(s.server.URL)

	created, err := resourceproviders.Create(
		context.Background(),
		client,
		resourceproviders.CreateOpts{
			Name: "compute-1",
			UUID: "resource-provider-1",
		},
	).Extract()
	s.Require().NoError(err)

	found, err := resourceproviders.Get(
		context.Background(),
		client,
		created.UUID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal("resource-provider-1", found.UUID)
	s.Assert().Equal("compute-1", found.Name)
	s.Assert().Equal("resource-provider-1", found.RootProviderUUID)
}

func (s *ResourceProviderSuite) TestListResourceProviders() {
	client := testhelper.ServiceClient(s.server.URL)

	_, err := resourceproviders.Create(
		context.Background(),
		client,
		resourceproviders.CreateOpts{
			Name: "compute-1",
			UUID: "resource-provider-1",
		},
	).Extract()
	s.Require().NoError(err)

	pages, err := resourceproviders.List(client, nil).AllPages(context.Background())
	s.Require().NoError(err)

	providers, err := resourceproviders.ExtractResourceProviders(pages)
	s.Require().NoError(err)

	s.Require().Len(providers, 1)
	s.Assert().Equal("resource-provider-1", providers[0].UUID)
	s.Assert().Equal("compute-1", providers[0].Name)
}

func (s *ResourceProviderSuite) TestDeleteResourceProvider() {
	client := testhelper.ServiceClient(s.server.URL)

	created, err := resourceproviders.Create(
		context.Background(),
		client,
		resourceproviders.CreateOpts{
			Name: "compute-1",
			UUID: "resource-provider-1",
		},
	).Extract()
	s.Require().NoError(err)

	err = resourceproviders.Delete(
		context.Background(),
		client,
		created.UUID,
	).ExtractErr()
	s.Require().NoError(err)

	pages, err := resourceproviders.List(client, nil).AllPages(context.Background())
	s.Require().NoError(err)

	providers, err := resourceproviders.ExtractResourceProviders(pages)
	s.Require().NoError(err)

	s.Assert().Empty(providers)
}
