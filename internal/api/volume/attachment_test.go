package volume_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/attachments"
	"github.com/stretchr/testify/suite"
)

type AttachmentSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestAttachmentSuite(t *testing.T) {
	suite.Run(t, new(AttachmentSuite))
}

func (s *AttachmentSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *AttachmentSuite) TearDownTest() {
	s.server.Close()
}

func (s *AttachmentSuite) TestCreateAndGetAttachment() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	created, err := attachments.Create(
		context.Background(),
		client,
		attachments.CreateOpts{
			VolumeUUID:   "volume-1",
			InstanceUUID: "server-1",
			Mode:         "rw",
		},
	).Extract()
	s.Require().NoError(err)

	found, err := attachments.Get(
		context.Background(),
		client,
		created.ID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal("volume-1", found.VolumeID)
	s.Assert().Equal("server-1", found.Instance)
	s.Assert().Equal("rw", found.AttachMode)
}

func (s *AttachmentSuite) TestListAttachments() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	_, err := attachments.Create(
		context.Background(),
		client,
		attachments.CreateOpts{
			VolumeUUID:   "volume-1",
			InstanceUUID: "server-1",
		},
	).Extract()
	s.Require().NoError(err)

	pages, err := attachments.List(client, nil).AllPages(context.Background())
	s.Require().NoError(err)

	found, err := attachments.ExtractAttachments(pages)
	s.Require().NoError(err)

	s.Require().Len(found, 1)
	s.Assert().Equal("volume-1", found[0].VolumeID)
}

func (s *AttachmentSuite) TestUpdateAndCompleteAttachment() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	created, err := attachments.Create(
		context.Background(),
		client,
		attachments.CreateOpts{
			VolumeUUID:   "volume-1",
			InstanceUUID: "server-1",
		},
	).Extract()
	s.Require().NoError(err)

	updated, err := attachments.Update(
		context.Background(),
		client,
		created.ID,
		attachments.UpdateOpts{
			Connector: map[string]any{
				"host": "compute-1",
			},
		},
	).Extract()
	s.Require().NoError(err)

	err = attachments.Complete(
		context.Background(),
		client,
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	found, err := attachments.Get(
		context.Background(),
		client,
		created.ID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal("attaching", updated.Status)
	s.Assert().Equal("attached", found.Status)
}

func (s *AttachmentSuite) TestDeleteAttachment() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	created, err := attachments.Create(
		context.Background(),
		client,
		attachments.CreateOpts{
			VolumeUUID:   "volume-1",
			InstanceUUID: "server-1",
		},
	).Extract()
	s.Require().NoError(err)

	err = attachments.Delete(
		context.Background(),
		client,
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	pages, err := attachments.List(client, nil).AllPages(context.Background())
	s.Require().NoError(err)

	found, err := attachments.ExtractAttachments(pages)
	s.Require().NoError(err)

	s.Assert().Empty(found)
}
