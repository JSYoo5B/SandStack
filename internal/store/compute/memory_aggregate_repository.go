package compute

import (
	"sync"

	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
)

type MemoryAggregateRepository struct {
	mu         sync.RWMutex
	ids        []int
	aggregates map[int]appcompute.Aggregate
}

func NewMemoryAggregateRepository() *MemoryAggregateRepository {
	return &MemoryAggregateRepository{
		aggregates: map[int]appcompute.Aggregate{},
	}
}

func (r *MemoryAggregateRepository) Create(
	aggregate appcompute.Aggregate,
) appcompute.Aggregate {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, aggregate.ID)
	r.aggregates[aggregate.ID] = aggregate

	return aggregate
}

func (r *MemoryAggregateRepository) List() []appcompute.Aggregate {
	r.mu.RLock()
	defer r.mu.RUnlock()

	aggregates := make([]appcompute.Aggregate, 0, len(r.ids))
	for _, id := range r.ids {
		aggregates = append(aggregates, r.aggregates[id])
	}

	return aggregates
}

func (r *MemoryAggregateRepository) Get(
	id int,
) (appcompute.Aggregate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	aggregate, ok := r.aggregates[id]
	if !ok {
		return appcompute.Aggregate{}, appcompute.ErrAggregateNotFound
	}

	return aggregate, nil
}

func (r *MemoryAggregateRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.aggregates[id]; !ok {
		return appcompute.ErrAggregateNotFound
	}
	delete(r.aggregates, id)
	for index, existingID := range r.ids {
		if existingID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryAggregateRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = nil
	r.aggregates = map[int]appcompute.Aggregate{}
}
