package image_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/image"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/images"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/members"
	"github.com/stretchr/testify/suite"
)

type MemberSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestMemberSuite(t *testing.T) {
	suite.Run(t, new(MemberSuite))
}

func (s *MemberSuite) SetupTest() {
	s.server = httptest.NewServer(
		image.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *MemberSuite) TearDownTest() {
	s.server.Close()
}

func (s *MemberSuite) TestCreateMemberThenListMembers() {
	image := s.createImage("ubuntu")

	created := s.createMember(image.ID, "project-1")
	list := s.listMembers(image.ID)

	s.Assert().Equal(image.ID, created.ImageID)
	s.Assert().Equal("project-1", created.MemberID)
	s.Assert().Equal("pending", created.Status)
	s.Require().Len(list, 1)
	s.Assert().Equal("project-1", list[0].MemberID)
}

func (s *MemberSuite) TestGetAndUpdateMember() {
	image := s.createImage("ubuntu")
	created := s.createMember(image.ID, "project-1")

	found, err := members.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		image.ID,
		created.MemberID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	updated, err := members.Update(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		image.ID,
		created.MemberID,
		members.UpdateOpts{Status: "accepted"},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(updated)

	s.Assert().Equal("project-1", found.MemberID)
	s.Assert().Equal("accepted", updated.Status)
}

func (s *MemberSuite) TestDeleteMember() {
	image := s.createImage("ubuntu")
	created := s.createMember(image.ID, "project-1")

	err := members.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		image.ID,
		created.MemberID,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listMembers(image.ID)

	s.Assert().Empty(list)
}

func (s *MemberSuite) listMembers(imageID string) []members.Member {
	pages, err := members.List(
		testhelper.ServiceClient(s.server.URL),
		imageID,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := members.ExtractMembers(pages)
	s.Require().NoError(err)

	return list
}

func (s *MemberSuite) createMember(
	imageID string,
	memberID string,
) *members.Member {
	created, err := members.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		imageID,
		memberID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *MemberSuite) createImage(name string) *images.Image {
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
