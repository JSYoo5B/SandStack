package identity_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/identity"
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
		identity.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *VersionSuite) TearDownTest() {
	s.server.Close()
}

func (s *VersionSuite) TestVersion() {
	response, err := http.Get(s.server.URL + "/")
	s.Require().NoError(err)
	defer response.Body.Close()

	var body versionResponse
	err = json.NewDecoder(response.Body).Decode(&body)
	s.Require().NoError(err)

	s.Assert().Equal(http.StatusOK, response.StatusCode)
	s.Assert().Equal("v3.14", body.Version.ID)
	s.Require().Len(body.Version.Links, 1)
	s.Assert().Equal(
		s.server.URL+"/identity/v3/",
		body.Version.Links[0].Href,
	)
}

func (s *VersionSuite) TestDiscovery() {
	handler := identity.NewHandler(testhelper.DefaultConfig())
	server := httptest.NewServer(
		handler.Discovery(),
	)
	defer server.Close()

	response, err := http.Get(server.URL)
	s.Require().NoError(err)
	defer response.Body.Close()

	var body versionsResponse
	err = json.NewDecoder(response.Body).Decode(&body)
	s.Require().NoError(err)

	s.Require().Len(body.Versions.Values, 1)
	version := body.Versions.Values[0]

	s.Assert().Equal(http.StatusOK, response.StatusCode)
	s.Assert().Equal("v3.14", version.ID)
	s.Require().Len(version.Links, 1)
	s.Assert().Equal(
		server.URL+"/identity/v3/",
		version.Links[0].Href,
	)
}

type versionsResponse struct {
	Versions versionValues `json:"versions"`
}

type versionValues struct {
	Values []versionDocument `json:"values"`
}

type versionResponse struct {
	Version versionDocument `json:"version"`
}

type versionDocument struct {
	ID    string        `json:"id"`
	Links []versionLink `json:"links"`
}

type versionLink struct {
	Href string `json:"href"`
}
