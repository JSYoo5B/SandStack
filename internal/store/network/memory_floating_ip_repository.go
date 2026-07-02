package network

import (
	"sync"

	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
)

type MemoryFloatingIPRepository struct {
	mu          sync.RWMutex
	ids         []string
	floatingIPs map[string]appnetwork.FloatingIP
}

func NewMemoryFloatingIPRepository() *MemoryFloatingIPRepository {
	return &MemoryFloatingIPRepository{
		ids:         []string{},
		floatingIPs: map[string]appnetwork.FloatingIP{},
	}
}

func (r *MemoryFloatingIPRepository) Create(
	floatingIP appnetwork.FloatingIP,
) appnetwork.FloatingIP {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, floatingIP.ID)
	r.floatingIPs[floatingIP.ID] = floatingIP

	return floatingIP
}

func (r *MemoryFloatingIPRepository) List() []appnetwork.FloatingIP {
	r.mu.RLock()
	defer r.mu.RUnlock()

	floatingIPs := make([]appnetwork.FloatingIP, 0, len(r.ids))
	for _, id := range r.ids {
		floatingIPs = append(floatingIPs, r.floatingIPs[id])
	}

	return floatingIPs
}

func (r *MemoryFloatingIPRepository) Get(
	id string,
) (appnetwork.FloatingIP, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	floatingIP, ok := r.floatingIPs[id]
	if !ok {
		return appnetwork.FloatingIP{}, appnetwork.ErrFloatingIPNotFound
	}

	return floatingIP, nil
}

func (r *MemoryFloatingIPRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.floatingIPs[id]; !ok {
		return appnetwork.ErrFloatingIPNotFound
	}

	delete(r.floatingIPs, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryFloatingIPRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.floatingIPs = map[string]appnetwork.FloatingIP{}
}
