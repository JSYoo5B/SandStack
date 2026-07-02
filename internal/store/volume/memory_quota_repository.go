package volume

import (
	"sync"

	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
)

type MemoryQuotaRepository struct {
	mu        sync.RWMutex
	quotaSets map[string]appvolume.QuotaSet
}

func NewMemoryQuotaRepository() *MemoryQuotaRepository {
	return &MemoryQuotaRepository{
		quotaSets: map[string]appvolume.QuotaSet{},
	}
}

func (r *MemoryQuotaRepository) Get(
	projectID string,
) (appvolume.QuotaSet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	quotaSet, ok := r.quotaSets[projectID]
	if !ok {
		return appvolume.QuotaSet{}, appvolume.ErrQuotaSetNotFound
	}

	return quotaSet, nil
}

func (r *MemoryQuotaRepository) Save(
	quotaSet appvolume.QuotaSet,
) appvolume.QuotaSet {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.quotaSets[quotaSet.ID] = quotaSet

	return quotaSet
}

func (r *MemoryQuotaRepository) Delete(projectID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.quotaSets, projectID)

	return nil
}

func (r *MemoryQuotaRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.quotaSets = map[string]appvolume.QuotaSet{}
}
