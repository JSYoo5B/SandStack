package fault

import "sync"

type Service struct {
	mu    sync.Mutex
	rules []Rule
	hits  map[string]int
	used  map[string]bool
}

func NewService() *Service {
	return &Service{
		rules: []Rule{},
		hits:  map[string]int{},
		used:  map[string]bool{},
	}
}

func (s *Service) Add(rule Rule) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.rules = append(s.rules, rule)
}

func (s *Service) List() []Rule {
	s.mu.Lock()
	defer s.mu.Unlock()

	rules := make([]Rule, 0, len(s.rules))
	rules = append(rules, s.rules...)

	return rules
}

func (s *Service) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.rules = []Rule{}
	s.hits = map[string]int{}
	s.used = map[string]bool{}
}

func (s *Service) Evaluate(operation Operation) Decision {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, rule := range s.rules {
		if !s.matches(rule, operation) {
			continue
		}
		if s.used[rule.ID] {
			continue
		}

		s.hits[rule.ID]++
		if !triggered(rule.Trigger, s.hits[rule.ID]) {
			continue
		}

		if rule.Trigger.Once {
			s.used[rule.ID] = true
		}

		return Decision{
			Matched:    true,
			HTTPStatus: rule.Behavior.HTTPStatus,
			Message:    rule.Behavior.Message,
		}
	}

	return Decision{}
}

func (s *Service) matches(rule Rule, operation Operation) bool {
	return rule.Enabled &&
		rule.Service == operation.Service &&
		rule.Operation == operation.Name
}

func triggered(trigger Trigger, hits int) bool {
	if trigger.Nth <= 0 {
		return true
	}

	return hits == trigger.Nth
}
