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
	pages, err := images.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := images.ExtractImages(pages)
	s.Require().NoError(err)

	s.Assert().Empty(list)
}
