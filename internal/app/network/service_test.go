package network_test

import (
	"testing"

	"github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storenetwork "github.com/JSYoo5B/SandStack/internal/store/network"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestCreateNetworkUsesInjectedIDGenerator() {
	service := newService(idgen.Fixed("network-id"))

	created := service.Create(network.CreateNetwork{Name: "private"})

	s.Assert().Equal("net-network-id", created.ID)
}

func (s *ServiceSuite) TestCreateSubnetUsesInjectedIDGenerator() {
	service := newService(idgen.Fixed("subnet-id"))

	created := service.CreateSubnet(network.CreateSubnet{
		NetworkID: "net-1",
		Name:      "private-subnet",
	})

	s.Assert().Equal("subnet-subnet-id", created.ID)
}

func (s *ServiceSuite) TestCreatePortUsesInjectedIDGenerator() {
	service := newService(idgen.Fixed("port-id"))

	created := service.CreatePort(network.CreatePort{
		NetworkID: "net-1",
		Name:      "private-port",
	})

	s.Assert().Equal("port-port-id", created.ID)
	s.Assert().Equal("fa:16:3e:port-id", created.MACAddress)
}

func (s *ServiceSuite) TestCreateSecurityGroupUsesInjectedIDGenerator() {
	service := newService(idgen.Fixed("security-group-id"))

	created := service.CreateSecurityGroup(network.CreateSecurityGroup{
		Name: "web",
	})

	s.Assert().Equal("sg-security-group-id", created.ID)
	s.Assert().Equal("web", created.Name)
	s.Assert().True(created.Stateful)
}

func (s *ServiceSuite) TestCreateSecurityGroupRuleUsesInjectedIDGenerator() {
	service := newService(idgen.Fixed("security-group-rule-id"))
	securityGroup := service.CreateSecurityGroup(network.CreateSecurityGroup{
		Name:      "web",
		ProjectID: "demo",
	})

	created, err := service.CreateSecurityGroupRule(
		network.CreateSecurityGroupRule{
			Direction:       "ingress",
			EtherType:       "IPv4",
			Protocol:        "tcp",
			PortRangeMin:    80,
			PortRangeMax:    80,
			RemoteIPPrefix:  "0.0.0.0/0",
			SecurityGroupID: securityGroup.ID,
		},
	)
	s.Require().NoError(err)

	s.Assert().Equal("sgr-security-group-rule-id", created.ID)
	s.Assert().Equal(securityGroup.ID, created.SecurityGroupID)
	s.Assert().Equal("demo", created.ProjectID)
}

func (s *ServiceSuite) TestCreateRouterUsesInjectedIDGenerator() {
	service := newService(idgen.Fixed("router-id"))

	created := service.CreateRouter(network.CreateRouter{
		Name: "edge",
	})

	s.Assert().Equal("router-router-id", created.ID)
	s.Assert().Equal("edge", created.Name)
	s.Assert().Equal("ACTIVE", created.Status)
	s.Assert().True(created.AdminStateUp)
}

func (s *ServiceSuite) TestResetClearsNetworkResources() {
	service := newService(idgen.Fixed("network-id"))
	created := service.Create(network.CreateNetwork{Name: "private"})
	service.CreateSubnet(network.CreateSubnet{NetworkID: created.ID})
	service.CreatePort(network.CreatePort{NetworkID: created.ID})
	securityGroup := service.CreateSecurityGroup(network.CreateSecurityGroup{
		Name: "default",
	})
	_, err := service.CreateSecurityGroupRule(network.CreateSecurityGroupRule{
		Direction:       "ingress",
		EtherType:       "IPv4",
		SecurityGroupID: securityGroup.ID,
	})
	s.Require().NoError(err)
	service.CreateRouter(network.CreateRouter{Name: "edge"})

	service.Reset()

	s.Assert().Empty(service.List())
	s.Assert().Empty(service.ListSubnets())
	s.Assert().Empty(service.ListPorts())
	s.Assert().Empty(service.ListSecurityGroups())
	s.Assert().Empty(service.ListSecurityGroupRules())
	s.Assert().Empty(service.ListRouters())
}

func newService(idGen idgen.Generator) *network.Service {
	return network.NewServiceWithRepositories(
		storenetwork.NewMemoryNetworkRepository(),
		storenetwork.NewMemorySubnetRepository(),
		storenetwork.NewMemoryPortRepository(),
		storenetwork.NewMemorySecurityGroupRepository(),
		storenetwork.NewMemorySecurityGroupRuleRepository(),
		storenetwork.NewMemoryRouterRepository(),
		idGen,
	)
}
