package network

import "errors"

var ErrSecurityGroupNotFound = errors.New("security group not found")

func (s *Service) CreateSecurityGroup(input CreateSecurityGroup) SecurityGroup {
	stateful := true
	if input.Stateful != nil {
		stateful = *input.Stateful
	}

	securityGroup := SecurityGroup{
		ID:          "sg-" + s.idGen.Hex(16),
		Name:        input.Name,
		Description: input.Description,
		Stateful:    stateful,
		TenantID:    input.ProjectID,
		ProjectID:   input.ProjectID,
		Rules:       []SecurityGroupRule{},
		Tags:        []string{},
	}

	return s.securityGroupRepository.Create(securityGroup)
}

func (s *Service) ListSecurityGroups() []SecurityGroup {
	return s.securityGroupRepository.List()
}

func (s *Service) GetSecurityGroup(id string) (SecurityGroup, error) {
	return s.securityGroupRepository.Get(id)
}

func (s *Service) DeleteSecurityGroup(id string) error {
	return s.securityGroupRepository.Delete(id)
}
