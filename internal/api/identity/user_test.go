package identity_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/users"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) SetupTest() {
	s.server = httptest.NewServer(
		identity.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *UserSuite) TearDownTest() {
	s.server.Close()
}

func (s *UserSuite) TestListUsers() {
	pages, err := users.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	listed, err := users.ExtractUsers(pages)
	s.Require().NoError(err)
	s.Require().Len(listed, 1)

	s.Assert().Equal("admin", listed[0].ID)
	s.Assert().Equal("admin", listed[0].Name)
	s.Assert().Equal("demo", listed[0].DefaultProjectID)
	s.Assert().True(listed[0].Enabled)
}

func (s *UserSuite) TestGetUser() {
	result := users.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		"admin",
	)
	user, err := result.Extract()
	s.Require().NoError(err)
	s.Require().NotNil(user)

	s.Assert().Equal("admin", user.ID)
	s.Assert().Equal("admin", user.Name)
	s.Assert().Equal("demo", user.DefaultProjectID)
	s.Assert().True(user.Enabled)
}
