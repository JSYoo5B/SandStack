package admin_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/admin"
	"github.com/stretchr/testify/suite"
)

type StatusSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestStatusSuite(t *testing.T) {
	suite.Run(t, new(StatusSuite))
}

func (s *StatusSuite) SetupTest() {
	s.server = httptest.NewServer(admin.NewRouter())
}

func (s *StatusSuite) TearDownTest() {
	s.server.Close()
}

func (s *StatusSuite) TestHealth() {
	s.assertStatusOK("/health")
}

func (s *StatusSuite) TestReady() {
	s.assertStatusOK("/ready")
}

func (s *StatusSuite) assertStatusOK(path string) {
	response, err := http.Get(s.server.URL + path)
	s.Require().NoError(err)
	defer response.Body.Close()

	var body map[string]string
	err = json.NewDecoder(response.Body).Decode(&body)
	s.Require().NoError(err)

	s.Assert().Equal(http.StatusOK, response.StatusCode)
	s.Assert().Equal(
		map[string]string{
			"status":  "ok",
			"service": "sandstack",
			"version": "0.1.0",
		},
		body,
	)
}
