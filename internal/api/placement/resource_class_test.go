package placement_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/placement"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/placement/v1/resourceclasses"
	"github.com/stretchr/testify/suite"
)

type ResourceClassSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestResourceClassSuite(t *testing.T) {
	suite.Run(t, new(ResourceClassSuite))
}

func (s *ResourceClassSuite) SetupTest() {
	s.server = httptest.NewServer(
		placement.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ResourceClassSuite) TearDownTest() {
	s.server.Close()
}

func (s *ResourceClassSuite) TestCreateAndGetResourceClass() {
	client := testhelper.ServiceClient(s.server.URL)

	err := resourceclasses.Create(
		context.Background(),
		client,
		resourceclasses.CreateOpts{Name: "CUSTOM_GPU"},
	).ExtractErr()
	s.Require().NoError(err)

	class, err := resourceclasses.Get(
		context.Background(),
		client,
		"CUSTOM_GPU",
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal("CUSTOM_GPU", class.Name)
}

func (s *ResourceClassSuite) TestUpdateResourceClass() {
	client := testhelper.ServiceClient(s.server.URL)

	err := resourceclasses.Update(
		context.Background(),
		client,
		"CUSTOM_FPGA",
	).ExtractErr()
	s.Require().NoError(err)

	pages, err := resourceclasses.List(client).AllPages(context.Background())
	s.Require().NoError(err)

	classes, err := resourceclasses.ExtractResourceClasses(pages)
	s.Require().NoError(err)

	s.Require().Len(classes, 1)
	s.Assert().Equal("CUSTOM_FPGA", classes[0].Name)
}

func (s *ResourceClassSuite) TestDeleteResourceClass() {
	client := testhelper.ServiceClient(s.server.URL)

	err := resourceclasses.Create(
		context.Background(),
		client,
		resourceclasses.CreateOpts{Name: "CUSTOM_GPU"},
	).ExtractErr()
	s.Require().NoError(err)

	err = resourceclasses.Delete(
		context.Background(),
		client,
		"CUSTOM_GPU",
	).ExtractErr()
	s.Require().NoError(err)

	pages, err := resourceclasses.List(client).AllPages(context.Background())
	s.Require().NoError(err)

	classes, err := resourceclasses.ExtractResourceClasses(pages)
	s.Require().NoError(err)

	s.Assert().Empty(classes)
}
