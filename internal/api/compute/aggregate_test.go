package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/aggregates"
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
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *AggregateSuite) TearDownTest() {
	s.server.Close()
}

func (s *AggregateSuite) TestCreateListGetAndDeleteAggregate() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	created, err := aggregates.Create(
		s.T().Context(),
		client,
		aggregates.CreateOpts{
			Name:             "az-one",
			AvailabilityZone: "nova",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	pages, err := aggregates.List(client).AllPages(s.T().Context())
	s.Require().NoError(err)
	listed, err := aggregates.ExtractAggregates(pages)
	s.Require().NoError(err)
	found, err := aggregates.Get(
		s.T().Context(),
		client,
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(1, created.ID)
	s.Assert().Equal("az-one", created.Name)
	s.Assert().Equal("nova", created.AvailabilityZone)
	s.Require().Len(listed, 1)
	s.Assert().Equal(created.ID, listed[0].ID)
	s.Assert().Equal(created.ID, found.ID)

	err = aggregates.Delete(
		s.T().Context(),
		client,
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)
}
