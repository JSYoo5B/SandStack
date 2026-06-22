package requestlog

import "sync"

type Service struct {
	mu      sync.RWMutex
	records []Record
}

func NewService() *Service {
	return &Service{records: []Record{}}
}

func (s *Service) Add(record Record) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.records = append(s.records, record)
}

func (s *Service) List() []Record {
	s.mu.RLock()
	defer s.mu.RUnlock()

	records := make([]Record, 0, len(s.records))
	records = append(records, s.records...)

	return records
}

func (s *Service) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.records = []Record{}
}
