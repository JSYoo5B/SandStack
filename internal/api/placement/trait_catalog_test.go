package placement_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/placement"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/placement/v1/traits"
	"github.com/stretchr/testify/suite"
)

type TraitCatalogSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestTraitCatalogSuite(t *testing.T) {
	suite.Run(t, new(TraitCatalogSuite))
}

func (s *TraitCatalogSuite) SetupTest() {
	s.server = httptest.NewServer(
		placement.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *TraitCatalogSuite) TearDownTest() {
	s.server.Close()
}

func (s *TraitCatalogSuite) TestCreateGetAndListTrait() {
	client := testhelper.ServiceClient(s.server.URL)

	err := traits.Create(
		context.Background(),
		client,
		"CUSTOM_GPU",
	).ExtractErr()
	s.Require().NoError(err)

	err = traits.Get(context.Background(), client, "CUSTOM_GPU").ExtractErr()
	s.Require().NoError(err)

	pages, err := traits.List(client, nil).AllPages(context.Background())
	s.Require().NoError(err)

	found, err := traits.ExtractTraits(pages)
	s.Require().NoError(err)

	s.Assert().Equal([]string{"CUSTOM_GPU"}, found)
}

func (s *TraitCatalogSuite) TestListTraitWithNameFilter() {
	client := testhelper.ServiceClient(s.server.URL)

	err := traits.Create(
		context.Background(),
		client,
		"CUSTOM_GPU",
	).ExtractErr()
	s.Require().NoError(err)
	err = traits.Create(
		context.Background(),
		client,
		"CUSTOM_FPGA",
	).ExtractErr()
	s.Require().NoError(err)

	pages, err := traits.List(
		client,
		traits.ListOpts{Name: "startswith:CUSTOM_G"},
	).AllPages(context.Background())
	s.Require().NoError(err)

	found, err := traits.ExtractTraits(pages)
	s.Require().NoError(err)

	s.Assert().Equal([]string{"CUSTOM_GPU"}, found)
}

func (s *TraitCatalogSuite) TestDeleteTrait() {
	client := testhelper.ServiceClient(s.server.URL)

	err := traits.Create(
		context.Background(),
		client,
		"CUSTOM_GPU",
	).ExtractErr()
	s.Require().NoError(err)

	err = traits.Delete(context.Background(), client, "CUSTOM_GPU").ExtractErr()
	s.Require().NoError(err)

	pages, err := traits.List(client, nil).AllPages(context.Background())
	s.Require().NoError(err)

	found, err := traits.ExtractTraits(pages)
	s.Require().NoError(err)

	s.Assert().Empty(found)
}
