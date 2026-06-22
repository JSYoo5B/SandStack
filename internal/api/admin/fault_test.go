package admin_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/admin"
	"github.com/stretchr/testify/suite"
)

type FaultSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestFaultSuite(t *testing.T) {
	suite.Run(t, new(FaultSuite))
}

func (s *FaultSuite) SetupTest() {
	s.server = httptest.NewServer(admin.NewRouter())
}

func (s *FaultSuite) TearDownTest() {
	s.server.Close()
}

func (s *FaultSuite) TestCreateThenListFaults() {
	created := s.createFault("rule-1")

	response, err := http.Get(s.server.URL + "/faults")
	s.Require().NoError(err)
	defer response.Body.Close()

	var body faultListResponse
	err = json.NewDecoder(response.Body).Decode(&body)
	s.Require().NoError(err)

	s.Assert().Equal(http.StatusCreated, created.StatusCode)
	s.Assert().Equal(http.StatusOK, response.StatusCode)
	s.Require().Len(body.Rules, 1)
	s.Assert().Equal("rule-1", body.Rules[0].ID)
	s.Assert().Equal("compute", body.Rules[0].Service)
	s.Assert().Equal("server.create", body.Rules[0].Operation)
	s.Assert().Equal(http.StatusServiceUnavailable, body.Rules[0].Behavior.HTTPStatus)
}

func (s *FaultSuite) TestDisableThenEnableFault() {
	response := s.createFault("rule-1")
	response.Body.Close()

	disabled := s.post("/faults/rule-1/disable", "")
	defer disabled.Body.Close()
	enabled := s.post("/faults/rule-1/enable", "")
	defer enabled.Body.Close()

	s.Assert().Equal(http.StatusNoContent, disabled.StatusCode)
	s.Assert().Equal(http.StatusNoContent, enabled.StatusCode)
}

func (s *FaultSuite) createFault(id string) *http.Response {
	body := `{
		"rule": {
			"id": "` + id + `",
			"service": "compute",
			"operation": "server.create",
			"behavior": {
				"http_status": 503,
				"message": "injected server create failure"
			},
			"trigger": {
				"nth": 3,
				"once": true
			}
		}
	}`

	return s.post("/faults", body)
}

func (s *FaultSuite) post(path string, body string) *http.Response {
	response, err := http.Post(
		s.server.URL+path,
		"application/json",
		bytes.NewBufferString(body),
	)
	s.Require().NoError(err)

	return response
}

type faultListResponse struct {
	Rules []faultRuleDocument `json:"rules"`
}

type faultRuleDocument struct {
	ID        string                `json:"id"`
	Service   string                `json:"service"`
	Operation string                `json:"operation"`
	Behavior  faultBehaviorDocument `json:"behavior"`
}

type faultBehaviorDocument struct {
	HTTPStatus int `json:"http_status"`
}
