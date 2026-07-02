package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/security/groups"
	"github.com/stretchr/testify/suite"
)

type SecurityGroupSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestSecurityGroupSuite(t *testing.T) {
	suite.Run(t, new(SecurityGroupSuite))
}

func (s *SecurityGroupSuite) SetupTest() {
	s.server = httptest.NewServer(
		network.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *SecurityGroupSuite) TearDownTest() {
	s.server.Close()
}

func (s *SecurityGroupSuite) TestListSecurityGroups() {
	list := s.listSecurityGroups()

	s.Assert().Empty(list)
}

func (s *SecurityGroupSuite) TestCreateSecurityGroupThenListSecurityGroups() {
	created := s.createSecurityGroup("web")

	list := s.listSecurityGroups()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal("web", created.Name)
	s.Assert().Equal("security group for web", created.Description)
	s.Assert().True(created.Stateful)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
	s.Assert().Equal("web", list[0].Name)
}

func (s *SecurityGroupSuite) TestGetSecurityGroup() {
	created := s.createSecurityGroup("web")

	found, err := groups.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.ID, found.ID)
	s.Assert().Equal("web", found.Name)
}

func (s *SecurityGroupSuite) TestDeleteSecurityGroup() {
	created := s.createSecurityGroup("web")

	err := groups.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listSecurityGroups()

	s.Assert().Empty(list)
}

func (s *SecurityGroupSuite) listSecurityGroups() []groups.SecGroup {
	pages, err := groups.List(
		testhelper.ServiceClient(s.server.URL),
		groups.ListOpts{},
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := groups.ExtractGroups(pages)
	s.Require().NoError(err)

	return list
}

func (s *SecurityGroupSuite) createSecurityGroup(name string) *groups.SecGroup {
	created, err := groups.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		groups.CreateOpts{
			Name:        name,
			Description: "security group for " + name,
			ProjectID:   "demo",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
