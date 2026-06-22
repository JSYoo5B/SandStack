package image_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/image"
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
		image.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *VersionSuite) TearDownTest() {
	s.server.Close()
}

func (s *VersionSuite) TestVersions() {
	response, err := http.Get(s.server.URL + "/")
	s.Require().NoError(err)
	defer response.Body.Close()

	var body versionsResponse
	err = json.NewDecoder(response.Body).Decode(&body)
	s.Require().NoError(err)

	s.Assert().Equal(http.StatusOK, response.StatusCode)
	s.Require().Len(body.Versions, 1)
	s.Assert().Equal("v2.0", body.Versions[0].ID)
	s.Assert().Equal("CURRENT", body.Versions[0].Status)
	s.Require().Len(body.Versions[0].Links, 1)
	s.Assert().Equal(
		s.server.URL+"/image/v2",
		body.Versions[0].Links[0].Href,
	)
}

type versionsResponse struct {
	Versions []versionDocument `json:"versions"`
}

type versionDocument struct {
	ID     string        `json:"id"`
	Status string        `json:"status"`
	Links  []versionLink `json:"links"`
}

type versionLink struct {
	Href string `json:"href"`
}
