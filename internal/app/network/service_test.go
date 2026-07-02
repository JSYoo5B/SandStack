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

func (s *ServiceSuite) TestResetClearsNetworkResources() {
	service := newService(idgen.Fixed("network-id"))
	created := service.Create(network.CreateNetwork{Name: "private"})
	service.CreateSubnet(network.CreateSubnet{NetworkID: created.ID})
	service.CreatePort(network.CreatePort{NetworkID: created.ID})

	service.Reset()

	s.Assert().Empty(service.List())
	s.Assert().Empty(service.ListSubnets())
	s.Assert().Empty(service.ListPorts())
}

func newService(idGen idgen.Generator) *network.Service {
	return network.NewServiceWithRepositories(
		storenetwork.NewMemoryNetworkRepository(),
		storenetwork.NewMemorySubnetRepository(),
		storenetwork.NewMemoryPortRepository(),
		idGen,
	)
}
