package identity_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"
	"github.com/stretchr/testify/suite"
)

type ProjectSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestProjectSuite(t *testing.T) {
	suite.Run(t, new(ProjectSuite))
}

func (s *ProjectSuite) SetupTest() {
	s.server = httptest.NewServer(
		identity.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *ProjectSuite) TearDownTest() {
	s.server.Close()
}

func (s *ProjectSuite) TestListProjects() {
	pages, err := projects.List(
		testhelper.ServiceClient(s.server.URL),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	listed, err := projects.ExtractProjects(pages)
	s.Require().NoError(err)
	s.Require().Len(listed, 1)

	s.Assert().Equal("demo", listed[0].ID)
	s.Assert().Equal("demo", listed[0].Name)
	s.Assert().Equal("default", listed[0].DomainID)
	s.Assert().True(listed[0].Enabled)
}

func (s *ProjectSuite) TestGetProject() {
	result := projects.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		"demo",
	)
	project, err := result.Extract()
	s.Require().NoError(err)
	s.Require().NotNil(project)

	s.Assert().Equal("demo", project.ID)
	s.Assert().Equal("demo", project.Name)
	s.Assert().Equal("default", project.DomainID)
	s.Assert().True(project.Enabled)
}
