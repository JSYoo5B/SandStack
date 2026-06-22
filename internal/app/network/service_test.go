package network_test

import (
	"testing"

	"github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestCreateNetworkUsesInjectedIDGenerator() {
	service := network.NewServiceWithIDGenerator(idgen.Fixed("network-id"))

	created := service.Create(network.CreateNetwork{Name: "private"})

	s.Assert().Equal("net-network-id", created.ID)
}

func (s *ServiceSuite) TestCreateSubnetUsesInjectedIDGenerator() {
	service := network.NewServiceWithIDGenerator(idgen.Fixed("subnet-id"))

	created := service.CreateSubnet(network.CreateSubnet{
		NetworkID: "net-1",
		Name:      "private-subnet",
	})

	s.Assert().Equal("subnet-subnet-id", created.ID)
}

func (s *ServiceSuite) TestCreatePortUsesInjectedIDGenerator() {
	service := network.NewServiceWithIDGenerator(idgen.Fixed("port-id"))

	created := service.CreatePort(network.CreatePort{
		NetworkID: "net-1",
		Name:      "private-port",
	})

	s.Assert().Equal("port-port-id", created.ID)
	s.Assert().Equal("fa:16:3e:port-id", created.MACAddress)
}
