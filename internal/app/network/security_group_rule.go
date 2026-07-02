package network

import "errors"

var ErrSecurityGroupRuleNotFound = errors.New("security group rule not found")

func (s *Service) CreateSecurityGroupRule(
	input CreateSecurityGroupRule,
) (SecurityGroupRule, error) {
	securityGroup, err := s.securityGroupRepository.Get(input.SecurityGroupID)
	if err != nil {
		return SecurityGroupRule{}, err
	}

	projectID := input.ProjectID
	if projectID == "" {
		projectID = securityGroup.ProjectID
	}

	rule := SecurityGroupRule{
		ID:                   "sgr-" + s.idGen.Hex(16),
		Direction:            input.Direction,
		Description:          input.Description,
		EtherType:            input.EtherType,
		Protocol:             input.Protocol,
		PortRangeMin:         input.PortRangeMin,
		PortRangeMax:         input.PortRangeMax,
		RemoteAddressGroupID: input.RemoteAddressGroupID,
		RemoteIPPrefix:       input.RemoteIPPrefix,
		RemoteGroupID:        input.RemoteGroupID,
		SecurityGroupID:      input.SecurityGroupID,
		TenantID:             projectID,
		ProjectID:            projectID,
	}

	return s.securityRuleRepository.Create(rule), nil
}

func (s *Service) ListSecurityGroupRules() []SecurityGroupRule {
	return s.securityRuleRepository.List()
}

func (s *Service) GetSecurityGroupRule(id string) (SecurityGroupRule, error) {
	return s.securityRuleRepository.Get(id)
}

func (s *Service) DeleteSecurityGroupRule(id string) error {
	return s.securityRuleRepository.Delete(id)
}
