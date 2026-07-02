package compute

import "errors"

var ErrServerGroupNotFound = errors.New("server group not found")

func (s *Service) CreateServerGroup(input CreateServerGroup) ServerGroup {
	policies := input.Policies
	policy := input.Policy
	if len(policies) == 0 && policy != "" {
		policies = []string{policy}
	}
	if policy == "" && len(policies) > 0 {
		policy = policies[0]
	}

	return s.groupRepository.Create(ServerGroup{
		ID:        "sg-" + s.idGen.Hex(16),
		Name:      input.Name,
		Policies:  policies,
		Members:   []string{},
		UserID:    "admin",
		ProjectID: "demo",
		Metadata:  map[string]any{},
		Policy:    policy,
	})
}

func (s *Service) ListServerGroups() []ServerGroup {
	return s.groupRepository.List()
}

func (s *Service) GetServerGroup(id string) (ServerGroup, error) {
	return s.groupRepository.Get(id)
}

func (s *Service) DeleteServerGroup(id string) error {
	return s.groupRepository.Delete(id)
}
