package image_test

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/image"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/imagedata"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/images"
	"github.com/stretchr/testify/suite"
)

type DataSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestDataSuite(t *testing.T) {
	suite.Run(t, new(DataSuite))
}

func (s *DataSuite) SetupTest() {
	s.server = httptest.NewServer(
		image.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *DataSuite) TearDownTest() {
	s.server.Close()
}

func (s *DataSuite) TestUploadAndDownloadImageData() {
	created := s.createImage("ubuntu")

	err := imagedata.Upload(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
		strings.NewReader("image-data"),
	).ExtractErr()
	s.Require().NoError(err)

	reader, err := imagedata.Download(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(reader)
	defer reader.Close()

	data, err := io.ReadAll(reader)
	s.Require().NoError(err)

	s.Assert().Equal("image-data", string(data))
}

func (s *DataSuite) createImage(name string) *images.Image {
	created, err := images.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
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
