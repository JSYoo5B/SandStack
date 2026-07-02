package image_test

import (
	"bytes"
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/image"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/imagedata"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/imageimport"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/images"
	"github.com/stretchr/testify/suite"
)

type ImportSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestImportSuite(t *testing.T) {
	suite.Run(t, new(ImportSuite))
}

func (s *ImportSuite) SetupTest() {
	s.server = httptest.NewServer(
		image.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ImportSuite) TearDownTest() {
	s.server.Close()
}

func (s *ImportSuite) TestGetImportInfo() {
	client := testhelper.ServiceClient(s.server.URL)

	info, err := imageimport.Get(context.Background(), client).Extract()
	s.Require().NoError(err)

	s.Assert().Contains(info.ImportMethods.Value, "glance-direct")
}

func (s *ImportSuite) TestStageAndImportImageData() {
	client := testhelper.ServiceClient(s.server.URL)

	created, err := images.Create(
		context.Background(),
		client,
		images.CreateOpts{Name: "ubuntu"},
	).Extract()
	s.Require().NoError(err)

	err = imagedata.Stage(
		context.Background(),
		client,
		created.ID,
		bytes.NewBufferString("image-bytes"),
	).ExtractErr()
	s.Require().NoError(err)

	staged, err := images.Get(context.Background(), client, created.ID).Extract()
	s.Require().NoError(err)
	s.Assert().Equal(images.ImageStatus("uploading"), staged.Status)

	err = imageimport.Create(
		context.Background(),
		client,
		created.ID,
		imageimport.CreateOpts{Name: imageimport.GlanceDirectMethod},
	).ExtractErr()
	s.Require().NoError(err)

	imported, err := images.Get(context.Background(), client, created.ID).Extract()
	s.Require().NoError(err)
	s.Assert().Equal(images.ImageStatusActive, imported.Status)
}
