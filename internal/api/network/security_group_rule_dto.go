package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type createSecurityGroupRuleRequest struct {
	SecurityGroupRule createSecurityGroupRuleDocument `json:"security_group_rule"`
}

type createSecurityGroupRuleDocument struct {
	Direction            string `json:"direction"`
	Description          string `json:"description"`
	EtherType            string `json:"ethertype"`
	Protocol             string `json:"protocol"`
	PortRangeMin         int    `json:"port_range_min"`
	PortRangeMax         int    `json:"port_range_max"`
	RemoteAddressGroupID string `json:"remote_address_group_id"`
	RemoteIPPrefix       string `json:"remote_ip_prefix"`
	RemoteGroupID        string `json:"remote_group_id"`
	SecurityGroupID      string `json:"security_group_id"`
	ProjectID            string `json:"project_id"`
	TenantID             string `json:"tenant_id"`
}

func (r createSecurityGroupRuleRequest) createSecurityGroupRule() appnetwork.CreateSecurityGroupRule {
	projectID := r.SecurityGroupRule.ProjectID
	if projectID == "" {
		projectID = r.SecurityGroupRule.TenantID
	}

	return appnetwork.CreateSecurityGroupRule{
		Direction:            r.SecurityGroupRule.Direction,
		Description:          r.SecurityGroupRule.Description,
		EtherType:            r.SecurityGroupRule.EtherType,
		Protocol:             r.SecurityGroupRule.Protocol,
		PortRangeMin:         r.SecurityGroupRule.PortRangeMin,
		PortRangeMax:         r.SecurityGroupRule.PortRangeMax,
		RemoteAddressGroupID: r.SecurityGroupRule.RemoteAddressGroupID,
		RemoteIPPrefix:       r.SecurityGroupRule.RemoteIPPrefix,
		RemoteGroupID:        r.SecurityGroupRule.RemoteGroupID,
		SecurityGroupID:      r.SecurityGroupRule.SecurityGroupID,
		ProjectID:            projectID,
	}
}

type securityGroupRuleListResponse struct {
	SecurityGroupRules []securityGroupRuleDocument `json:"security_group_rules"`
}

type securityGroupRuleResponse struct {
	SecurityGroupRule securityGroupRuleDocument `json:"security_group_rule"`
}

func toSecurityGroupRuleDocument(
	rule appnetwork.SecurityGroupRule,
) securityGroupRuleDocument {
	return securityGroupRuleDocument{
		ID:                   rule.ID,
		Direction:            rule.Direction,
		Description:          rule.Description,
		EtherType:            rule.EtherType,
		Protocol:             rule.Protocol,
		PortRangeMin:         rule.PortRangeMin,
		PortRangeMax:         rule.PortRangeMax,
		RemoteAddressGroupID: rule.RemoteAddressGroupID,
		RemoteIPPrefix:       rule.RemoteIPPrefix,
		RemoteGroupID:        rule.RemoteGroupID,
		SecurityGroupID:      rule.SecurityGroupID,
		TenantID:             rule.TenantID,
		ProjectID:            rule.ProjectID,
	}
}
