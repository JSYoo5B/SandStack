package compute_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/stretchr/testify/suite"
)

type VersionSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestVersionSuite(t *testing.T) {
	suite.Run(t, new(VersionSuite))
}

func (s *VersionSuite) SetupTest() {
	s.server = httptest.NewServer(
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *VersionSuite) TearDownTest() {
	s.server.Close()
}

func (s *VersionSuite) TestVersion() {
	response, err := http.Get(s.server.URL + "/demo")
	s.Require().NoError(err)
	defer response.Body.Close()

	var body versionResponse
	err = json.NewDecoder(response.Body).Decode(&body)
	s.Require().NoError(err)

	s.Assert().Equal(http.StatusOK, response.StatusCode)
	s.Assert().Equal("v2.1", body.Version.ID)
	s.Assert().Equal("CURRENT", body.Version.Status)
	s.Require().Len(body.Version.Links, 1)
	s.Assert().Equal(
		s.server.URL+"/compute/v2.1/demo",
		body.Version.Links[0].Href,
	)
}

type versionResponse struct {
	Version versionDocument `json:"version"`
}

type versionDocument struct {
	ID     string        `json:"id"`
	Status string        `json:"status"`
	Links  []versionLink `json:"links"`
}

type versionLink struct {
	Href string `json:"href"`
}
