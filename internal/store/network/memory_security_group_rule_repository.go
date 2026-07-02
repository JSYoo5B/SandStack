package network

import (
	"sync"

	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
)

type MemorySecurityGroupRuleRepository struct {
	mu    sync.RWMutex
	ids   []string
	rules map[string]appnetwork.SecurityGroupRule
}

func NewMemorySecurityGroupRuleRepository() *MemorySecurityGroupRuleRepository {
	return &MemorySecurityGroupRuleRepository{
		ids:   []string{},
		rules: map[string]appnetwork.SecurityGroupRule{},
	}
}

func (r *MemorySecurityGroupRuleRepository) Create(
	rule appnetwork.SecurityGroupRule,
) appnetwork.SecurityGroupRule {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, rule.ID)
	r.rules[rule.ID] = rule

	return rule
}

func (r *MemorySecurityGroupRuleRepository) List() []appnetwork.SecurityGroupRule {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rules := make([]appnetwork.SecurityGroupRule, 0, len(r.ids))
	for _, id := range r.ids {
		rules = append(rules, r.rules[id])
	}

	return rules
}

func (r *MemorySecurityGroupRuleRepository) Get(
	id string,
) (appnetwork.SecurityGroupRule, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rule, ok := r.rules[id]
	if !ok {
		return appnetwork.SecurityGroupRule{}, appnetwork.ErrSecurityGroupRuleNotFound
	}

	return rule, nil
}

func (r *MemorySecurityGroupRuleRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.rules[id]; !ok {
		return appnetwork.ErrSecurityGroupRuleNotFound
	}

	delete(r.rules, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemorySecurityGroupRuleRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.rules = map[string]appnetwork.SecurityGroupRule{}
}
