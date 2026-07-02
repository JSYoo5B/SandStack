package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type createSecurityGroupRequest struct {
	SecurityGroup createSecurityGroupDocument `json:"security_group"`
}

type createSecurityGroupDocument struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stateful    *bool  `json:"stateful"`
	ProjectID   string `json:"project_id"`
	TenantID    string `json:"tenant_id"`
}

func (r createSecurityGroupRequest) createSecurityGroup() appnetwork.CreateSecurityGroup {
	projectID := r.SecurityGroup.ProjectID
	if projectID == "" {
		projectID = r.SecurityGroup.TenantID
	}

	return appnetwork.CreateSecurityGroup{
		Name:        r.SecurityGroup.Name,
		Description: r.SecurityGroup.Description,
		Stateful:    r.SecurityGroup.Stateful,
		ProjectID:   projectID,
	}
}

type securityGroupListResponse struct {
	SecurityGroups []securityGroupDocument `json:"security_groups"`
}

type securityGroupResponse struct {
	SecurityGroup securityGroupDocument `json:"security_group"`
}

type securityGroupDocument struct {
	ID          string                      `json:"id"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Stateful    bool                        `json:"stateful"`
	TenantID    string                      `json:"tenant_id"`
	ProjectID   string                      `json:"project_id"`
	Rules       []securityGroupRuleDocument `json:"security_group_rules"`
	Tags        []string                    `json:"tags"`
}

type securityGroupRuleDocument struct {
	ID                   string `json:"id"`
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
	TenantID             string `json:"tenant_id"`
	ProjectID            string `json:"project_id"`
}

func toSecurityGroupDocuments(
	securityGroups []appnetwork.SecurityGroup,
) []securityGroupDocument {
	documents := make([]securityGroupDocument, 0, len(securityGroups))
	for _, securityGroup := range securityGroups {
		documents = append(documents, toSecurityGroupDocument(securityGroup))
	}

	return documents
}

func toSecurityGroupDocument(
	securityGroup appnetwork.SecurityGroup,
) securityGroupDocument {
	return securityGroupDocument{
		ID:          securityGroup.ID,
		Name:        securityGroup.Name,
		Description: securityGroup.Description,
		Stateful:    securityGroup.Stateful,
		TenantID:    securityGroup.TenantID,
		ProjectID:   securityGroup.ProjectID,
		Rules:       toSecurityGroupRuleDocuments(securityGroup.Rules),
		Tags:        securityGroup.Tags,
	}
}

func toSecurityGroupRuleDocuments(
	rules []appnetwork.SecurityGroupRule,
) []securityGroupRuleDocument {
	documents := make([]securityGroupRuleDocument, 0, len(rules))
	for _, rule := range rules {
		documents = append(documents, securityGroupRuleDocument{
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
		})
	}

	return documents
}
