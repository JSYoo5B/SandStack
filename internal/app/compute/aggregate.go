package compute

import "errors"

var ErrAggregateNotFound = errors.New("aggregate not found")

const aggregateTimestampFormat = "2006-01-02T15:04:05.000000"

func (s *Service) CreateAggregate(input CreateAggregate) Aggregate {
	now := s.clock.Now().UTC().Format(aggregateTimestampFormat)
	id := len(s.aggregateRepository.List()) + 1

	return s.aggregateRepository.Create(Aggregate{
		ID:               id,
		Name:             input.Name,
		AvailabilityZone: input.AvailabilityZone,
		Hosts:            []string{},
		Metadata:         map[string]string{},
		CreatedAt:        now,
		UpdatedAt:        now,
		UUID:             "aggregate-" + s.idGen.Hex(16),
	})
}

func (s *Service) ListAggregates() []Aggregate {
	return s.aggregateRepository.List()
}

func (s *Service) GetAggregate(id int) (Aggregate, error) {
	return s.aggregateRepository.Get(id)
}

func (s *Service) DeleteAggregate(id int) error {
	return s.aggregateRepository.Delete(id)
}
