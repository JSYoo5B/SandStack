package network

import (
	"sync"

	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
)

type MemoryRouterRepository struct {
	mu      sync.RWMutex
	ids     []string
	routers map[string]appnetwork.Router
}

func NewMemoryRouterRepository() *MemoryRouterRepository {
	return &MemoryRouterRepository{
		ids:     []string{},
		routers: map[string]appnetwork.Router{},
	}
}

func (r *MemoryRouterRepository) Create(router appnetwork.Router) appnetwork.Router {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, router.ID)
	r.routers[router.ID] = router

	return router
}

func (r *MemoryRouterRepository) List() []appnetwork.Router {
	r.mu.RLock()
	defer r.mu.RUnlock()

	routers := make([]appnetwork.Router, 0, len(r.ids))
	for _, id := range r.ids {
		routers = append(routers, r.routers[id])
	}

	return routers
}

func (r *MemoryRouterRepository) Get(id string) (appnetwork.Router, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	router, ok := r.routers[id]
	if !ok {
		return appnetwork.Router{}, appnetwork.ErrRouterNotFound
	}

	return router, nil
}

func (r *MemoryRouterRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.routers[id]; !ok {
		return appnetwork.ErrRouterNotFound
	}

	delete(r.routers, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryRouterRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.routers = map[string]appnetwork.Router{}
}
