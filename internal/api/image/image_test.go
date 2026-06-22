package image_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/image"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/images"
	"github.com/stretchr/testify/suite"
)

type ImageSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestImageSuite(t *testing.T) {
	suite.Run(t, new(ImageSuite))
}

func (s *ImageSuite) SetupTest() {
	s.server = httptest.NewServer(
		image.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ImageSuite) TearDownTest() {
	s.server.Close()
}

func (s *ImageSuite) TestListImages() {
	list := s.listImages()

	s.Assert().Empty(list)
}

func (s *ImageSuite) TestCreateImageThenListImages() {
	created, err := images.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		images.CreateOpts{
			Name:            "ubuntu",
			ContainerFormat: "bare",
			DiskFormat:      "qcow2",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	list := s.listImages()

	s.Assert().NotEmpty(created.ID)
	s.Assert().Equal("ubuntu", created.Name)
	s.Assert().Equal(images.ImageStatus("queued"), created.Status)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
	s.Assert().Equal("ubuntu", list[0].Name)
}

func (s *ImageSuite) listImages() []images.Image {
	pages, err := images.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := images.ExtractImages(pages)
	s.Require().NoError(err)

	return list
}
