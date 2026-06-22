package admin

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/app/fault"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type createFaultRequest struct {
	Rule createFaultRuleDocument `json:"rule"`
}

type faultListResponse struct {
	Rules []faultRuleDocument `json:"rules"`
}

type faultResponse struct {
	Rule faultRuleDocument `json:"rule"`
}

type createFaultRuleDocument struct {
	ID        string                `json:"id"`
	Enabled   *bool                 `json:"enabled"`
	Service   string                `json:"service"`
	Operation string                `json:"operation"`
	Behavior  faultBehaviorDocument `json:"behavior"`
	Trigger   faultTriggerDocument  `json:"trigger"`
}

type faultRuleDocument struct {
	ID        string                `json:"id"`
	Enabled   bool                  `json:"enabled"`
	Service   string                `json:"service"`
	Operation string                `json:"operation"`
	Behavior  faultBehaviorDocument `json:"behavior"`
	Trigger   faultTriggerDocument  `json:"trigger"`
}

type faultBehaviorDocument struct {
	HTTPStatus int    `json:"http_status"`
	Message    string `json:"message"`
}

type faultTriggerDocument struct {
	Nth  int  `json:"nth"`
	Once bool `json:"once"`
}

func (r createFaultRequest) rule() fault.Rule {
	id := r.Rule.ID
	if id == "" {
		id = "fault-" + idgen.RandomHex(8)
	}

	enabled := true
	if r.Rule.Enabled != nil {
		enabled = *r.Rule.Enabled
	}

	status := r.Rule.Behavior.HTTPStatus
	if status == 0 {
		status = http.StatusServiceUnavailable
	}

	return fault.Rule{
		ID:        id,
		Enabled:   enabled,
		Service:   r.Rule.Service,
		Operation: r.Rule.Operation,
		Behavior: fault.Behavior{
			HTTPStatus: status,
			Message:    r.Rule.Behavior.Message,
		},
		Trigger: fault.Trigger{
			Nth:  r.Rule.Trigger.Nth,
			Once: r.Rule.Trigger.Once,
		},
	}
}

func toFaultRuleDocuments(rules []fault.Rule) []faultRuleDocument {
	documents := make([]faultRuleDocument, 0, len(rules))
	for _, rule := range rules {
		documents = append(documents, toFaultRuleDocument(rule))
	}

	return documents
}

func toFaultRuleDocument(rule fault.Rule) faultRuleDocument {
	return faultRuleDocument{
		ID:        rule.ID,
		Enabled:   rule.Enabled,
		Service:   rule.Service,
		Operation: rule.Operation,
		Behavior: faultBehaviorDocument{
			HTTPStatus: rule.Behavior.HTTPStatus,
			Message:    rule.Behavior.Message,
		},
		Trigger: faultTriggerDocument{
			Nth:  rule.Trigger.Nth,
			Once: rule.Trigger.Once,
		},
	}
}
