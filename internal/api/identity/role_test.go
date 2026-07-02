package identity_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"
	"github.com/stretchr/testify/suite"
)

type RoleSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestRoleSuite(t *testing.T) {
	suite.Run(t, new(RoleSuite))
}

func (s *RoleSuite) SetupTest() {
	s.server = httptest.NewServer(
		identity.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *RoleSuite) TearDownTest() {
	s.server.Close()
}

func (s *RoleSuite) TestListRoles() {
	pages, err := roles.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	listed, err := roles.ExtractRoles(pages)
	s.Require().NoError(err)
	s.Require().Len(listed, 1)

	s.Assert().Equal("admin", listed[0].ID)
	s.Assert().Equal("admin", listed[0].Name)
	s.Assert().Equal("default", listed[0].DomainID)
}

func (s *RoleSuite) TestGetRole() {
	result := roles.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		"admin",
	)
	role, err := result.Extract()
	s.Require().NoError(err)
	s.Require().NotNil(role)

	s.Assert().Equal("admin", role.ID)
	s.Assert().Equal("admin", role.Name)
	s.Assert().Equal("default", role.DomainID)
}
