package network_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/security/rules"
	"github.com/stretchr/testify/suite"
)

type SecurityGroupRuleSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestSecurityGroupRuleSuite(t *testing.T) {
	suite.Run(t, new(SecurityGroupRuleSuite))
}

func (s *SecurityGroupRuleSuite) SetupTest() {
	s.server = httptest.NewServer(
		network.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *SecurityGroupRuleSuite) TearDownTest() {
	s.server.Close()
}

func (s *SecurityGroupRuleSuite) TestListSecurityGroupRules() {
	list := s.listSecurityGroupRules()

	s.Assert().Empty(list)
}

func (s *SecurityGroupRuleSuite) TestCreateSecurityGroupRuleThenListRules() {
	securityGroup := s.createSecurityGroup("web")
	created := s.createSecurityGroupRule(securityGroup.ID)

	list := s.listSecurityGroupRules()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal(securityGroup.ID, created.SecGroupID)
	s.Assert().Equal("ingress", created.Direction)
	s.Assert().Equal("IPv4", created.EtherType)
	s.Assert().Equal("tcp", created.Protocol)
	s.Assert().Equal("0.0.0.0/0", created.RemoteIPPrefix)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
}

func (s *SecurityGroupRuleSuite) TestGetSecurityGroupRule() {
	securityGroup := s.createSecurityGroup("web")
	created := s.createSecurityGroupRule(securityGroup.ID)

	found, err := rules.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.ID, found.ID)
	s.Assert().Equal(securityGroup.ID, found.SecGroupID)
}

func (s *SecurityGroupRuleSuite) TestDeleteSecurityGroupRule() {
	securityGroup := s.createSecurityGroup("web")
	created := s.createSecurityGroupRule(securityGroup.ID)

	err := rules.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listSecurityGroupRules()

	s.Assert().Empty(list)
}

func (s *SecurityGroupRuleSuite) listSecurityGroupRules() []rules.SecGroupRule {
	pages, err := rules.List(
		testhelper.ServiceClient(s.server.URL),
		rules.ListOpts{},
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := rules.ExtractRules(pages)
	s.Require().NoError(err)

	return list
}

func (s *SecurityGroupRuleSuite) createSecurityGroup(name string) *groups.SecGroup {
	created, err := groups.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		groups.CreateOpts{
			Name:      name,
			ProjectID: "demo",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *SecurityGroupRuleSuite) createSecurityGroupRule(
	securityGroupID string,
) *rules.SecGroupRule {
	created, err := rules.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		rules.CreateOpts{
			Direction:      rules.DirIngress,
			EtherType:      rules.EtherType4,
			Protocol:       rules.ProtocolTCP,
			PortRangeMin:   80,
			PortRangeMax:   80,
			RemoteIPPrefix: "0.0.0.0/0",
			SecGroupID:     securityGroupID,
			ProjectID:      "demo",
			Description:    "allow web",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
