package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/flavors"
	"github.com/stretchr/testify/suite"
)

type FlavorExtraSpecSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestFlavorExtraSpecSuite(t *testing.T) {
	suite.Run(t, new(FlavorExtraSpecSuite))
}

func (s *FlavorExtraSpecSuite) SetupTest() {
	s.server = httptest.NewServer(
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *FlavorExtraSpecSuite) TearDownTest() {
	s.server.Close()
}

func (s *FlavorExtraSpecSuite) TestManageFlavorExtraSpecs() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	created, err := flavors.CreateExtraSpecs(
		s.T().Context(),
		client,
		"1",
		flavors.ExtraSpecsOpts{
			"hw:cpu_policy": "shared",
		},
	).Extract()
	s.Require().NoError(err)
	s.Assert().Equal("shared", created["hw:cpu_policy"])

	list, err := flavors.ListExtraSpecs(
		s.T().Context(),
		client,
		"1",
	).Extract()
	s.Require().NoError(err)
	s.Assert().Equal("shared", list["hw:cpu_policy"])

	found, err := flavors.GetExtraSpec(
		s.T().Context(),
		client,
		"1",
		"hw:cpu_policy",
	).Extract()
	s.Require().NoError(err)
	s.Assert().Equal("shared", found["hw:cpu_policy"])

	updated, err := flavors.UpdateExtraSpec(
		s.T().Context(),
		client,
		"1",
		flavors.ExtraSpecsOpts{
			"hw:cpu_policy": "dedicated",
		},
	).Extract()
	s.Require().NoError(err)
	s.Assert().Equal("dedicated", updated["hw:cpu_policy"])

	err = flavors.DeleteExtraSpec(
		s.T().Context(),
		client,
		"1",
		"hw:cpu_policy",
	).ExtractErr()
	s.Require().NoError(err)
}
