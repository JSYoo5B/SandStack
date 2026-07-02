package volume_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumetypes"
	"github.com/stretchr/testify/suite"
)

type ExtraSpecSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestExtraSpecSuite(t *testing.T) {
	suite.Run(t, new(ExtraSpecSuite))
}

func (s *ExtraSpecSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ExtraSpecSuite) TearDownTest() {
	s.server.Close()
}

func (s *ExtraSpecSuite) TestCreateListAndGetExtraSpecs() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	created, err := volumetypes.CreateExtraSpecs(
		context.Background(),
		client,
		"default",
		volumetypes.ExtraSpecsOpts{
			"capabilities": "gpu",
		},
	).Extract()
	s.Require().NoError(err)

	listed, err := volumetypes.ListExtraSpecs(
		context.Background(),
		client,
		"default",
	).Extract()
	s.Require().NoError(err)

	found, err := volumetypes.GetExtraSpec(
		context.Background(),
		client,
		"default",
		"capabilities",
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal("gpu", created["capabilities"])
	s.Assert().Equal("gpu", listed["capabilities"])
	s.Assert().Equal("gpu", found["capabilities"])
}

func (s *ExtraSpecSuite) TestUpdateExtraSpec() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	updated, err := volumetypes.UpdateExtraSpec(
		context.Background(),
		client,
		"default",
		volumetypes.ExtraSpecsOpts{
			"capabilities": "ssd",
		},
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal("ssd", updated["capabilities"])
}

func (s *ExtraSpecSuite) TestDeleteExtraSpec() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	_, err := volumetypes.CreateExtraSpecs(
		context.Background(),
		client,
		"default",
		volumetypes.ExtraSpecsOpts{
			"capabilities": "gpu",
		},
	).Extract()
	s.Require().NoError(err)

	err = volumetypes.DeleteExtraSpec(
		context.Background(),
		client,
		"default",
		"capabilities",
	).ExtractErr()
	s.Require().NoError(err)

	listed, err := volumetypes.ListExtraSpecs(
		context.Background(),
		client,
		"default",
	).Extract()
	s.Require().NoError(err)

	s.Assert().Empty(listed)
}
