package fault_test

import (
	"net/http"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/app/fault"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestEvaluateMatchesEnabledRule() {
	service := fault.NewService()
	service.Add(serverCreateRule(fault.Trigger{}))

	decision := service.Evaluate(serverCreateOperation())

	s.Assert().True(decision.Matched)
	s.Assert().Equal(http.StatusServiceUnavailable, decision.HTTPStatus)
	s.Assert().Equal("injected server create failure", decision.Message)
}

func (s *ServiceSuite) TestEvaluateIgnoresDisabledRule() {
	service := fault.NewService()
	rule := serverCreateRule(fault.Trigger{})
	rule.Enabled = false
	service.Add(rule)

	decision := service.Evaluate(serverCreateOperation())

	s.Assert().False(decision.Matched)
}

func (s *ServiceSuite) TestEvaluateMatchesNthAttempt() {
	service := fault.NewService()
	service.Add(serverCreateRule(fault.Trigger{Nth: 3}))

	first := service.Evaluate(serverCreateOperation())
	second := service.Evaluate(serverCreateOperation())
	third := service.Evaluate(serverCreateOperation())

	s.Assert().False(first.Matched)
	s.Assert().False(second.Matched)
	s.Assert().True(third.Matched)
}

func (s *ServiceSuite) TestEvaluateMatchesOnce() {
	service := fault.NewService()
	service.Add(serverCreateRule(fault.Trigger{Once: true}))

	first := service.Evaluate(serverCreateOperation())
	second := service.Evaluate(serverCreateOperation())

	s.Assert().True(first.Matched)
	s.Assert().False(second.Matched)
}

func (s *ServiceSuite) TestDisableRule() {
	service := fault.NewService()
	service.Add(serverCreateRule(fault.Trigger{}))

	err := service.Disable("rule-1")
	s.Require().NoError(err)

	decision := service.Evaluate(serverCreateOperation())
	s.Assert().False(decision.Matched)
}

func (s *ServiceSuite) TestEnableRule() {
	service := fault.NewService()
	rule := serverCreateRule(fault.Trigger{})
	rule.Enabled = false
	service.Add(rule)

	err := service.Enable("rule-1")
	s.Require().NoError(err)

	decision := service.Evaluate(serverCreateOperation())
	s.Assert().True(decision.Matched)
}

func serverCreateRule(trigger fault.Trigger) fault.Rule {
	return fault.Rule{
		ID:        "rule-1",
		Enabled:   true,
		Service:   "compute",
		Operation: "server.create",
		Behavior: fault.Behavior{
			HTTPStatus: http.StatusServiceUnavailable,
			Message:    "injected server create failure",
		},
		Trigger: trigger,
	}
}

func serverCreateOperation() fault.Operation {
	return fault.Operation{
		Service: "compute",
		Name:    "server.create",
	}
}
