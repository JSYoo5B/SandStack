package volume_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/stretchr/testify/suite"
)

type VolumeActionSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestVolumeActionSuite(t *testing.T) {
	suite.Run(t, new(VolumeActionSuite))
}

func (s *VolumeActionSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *VolumeActionSuite) TearDownTest() {
	s.server.Close()
}

func (s *VolumeActionSuite) TestAttachAndDetachVolume() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	created := s.createVolume(client)

	err := volumes.Attach(
		context.Background(),
		client,
		created.ID,
		volumes.AttachOpts{
			InstanceUUID: "server-1",
			Mode:         volumes.ReadWrite,
		},
	).ExtractErr()
	s.Require().NoError(err)

	attached, err := volumes.Get(context.Background(), client, created.ID).Extract()
	s.Require().NoError(err)
	s.Assert().Equal("in-use", attached.Status)

	err = volumes.Detach(
		context.Background(),
		client,
		created.ID,
		volumes.DetachOpts{},
	).ExtractErr()
	s.Require().NoError(err)

	detached, err := volumes.Get(context.Background(), client, created.ID).Extract()
	s.Require().NoError(err)
	s.Assert().Equal("available", detached.Status)
}

func (s *VolumeActionSuite) TestReserveAndUnreserveVolume() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	created := s.createVolume(client)

	err := volumes.Reserve(context.Background(), client, created.ID).ExtractErr()
	s.Require().NoError(err)

	reserved, err := volumes.Get(context.Background(), client, created.ID).Extract()
	s.Require().NoError(err)
	s.Assert().Equal("reserved", reserved.Status)

	err = volumes.Unreserve(context.Background(), client, created.ID).ExtractErr()
	s.Require().NoError(err)

	unreserved, err := volumes.Get(context.Background(), client, created.ID).Extract()
	s.Require().NoError(err)
	s.Assert().Equal("available", unreserved.Status)
}

func (s *VolumeActionSuite) TestExtendAndResetStatusVolume() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	created := s.createVolume(client)

	err := volumes.ExtendSize(
		context.Background(),
		client,
		created.ID,
		volumes.ExtendSizeOpts{NewSize: 2},
	).ExtractErr()
	s.Require().NoError(err)

	extended, err := volumes.Get(context.Background(), client, created.ID).Extract()
	s.Require().NoError(err)
	s.Assert().Equal(2, extended.Size)

	err = volumes.ResetStatus(
		context.Background(),
		client,
		created.ID,
		volumes.ResetStatusOpts{Status: "error"},
	).ExtractErr()
	s.Require().NoError(err)

	reset, err := volumes.Get(context.Background(), client, created.ID).Extract()
	s.Require().NoError(err)
	s.Assert().Equal("error", reset.Status)
}

func (s *VolumeActionSuite) createVolume(
	client *gophercloud.ServiceClient,
) *volumes.Volume {
	created, err := volumes.Create(
		context.Background(),
		client,
		volumes.CreateOpts{
			Size: 1,
			Name: "database",
		},
		nil,
	).Extract()
	s.Require().NoError(err)

	return created
}
