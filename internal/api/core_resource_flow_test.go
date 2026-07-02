package api_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/images"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	"github.com/stretchr/testify/suite"
)

type CoreResourceFlowSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestCoreResourceFlowSuite(t *testing.T) {
	suite.Run(t, new(CoreResourceFlowSuite))
}

func (s *CoreResourceFlowSuite) SetupTest() {
	s.server = httptest.NewServer(api.NewRouter(testhelper.DefaultConfig()))
}

func (s *CoreResourceFlowSuite) TearDownTest() {
	s.server.Close()
}

func (s *CoreResourceFlowSuite) TestCreateCoreResources() {
	image := s.createImage("ubuntu")
	network := s.createNetwork("private")
	server := s.createServer("web", image.ID, network.ID)
	volume := s.createVolume("database")

	s.Assert().NotEmpty(image.ID)
	s.Assert().Equal("ubuntu", image.Name)
	s.Assert().NotEmpty(network.ID)
	s.Assert().Equal("private", network.Name)
	s.Assert().NotEmpty(server.ID)
	s.Assert().Equal("web", server.Name)
	s.Assert().NotEmpty(volume.ID)
	s.Assert().Equal("database", volume.Name)
}

func (s *CoreResourceFlowSuite) createImage(name string) *images.Image {
	created, err := images.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/image/v2"),
		images.CreateOpts{
			Name:            name,
			ContainerFormat: "bare",
			DiskFormat:      "qcow2",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *CoreResourceFlowSuite) createNetwork(
	name string,
) *networks.Network {
	created, err := networks.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/network/v2.0"),
		networks.CreateOpts{
			Name:      name,
			ProjectID: "demo",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *CoreResourceFlowSuite) createServer(
	name string,
	imageID string,
	networkID string,
) *servers.Server {
	created, err := servers.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/compute/v2.1/demo"),
		servers.CreateOpts{
			Name:      name,
			ImageRef:  imageID,
			FlavorRef: "1",
			Networks: []servers.Network{
				{UUID: networkID},
			},
		},
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *CoreResourceFlowSuite) createVolume(name string) *volumes.Volume {
	created, err := volumes.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/volume/v3/demo"),
		volumes.CreateOpts{
			Size: 1,
			Name: name,
		},
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
