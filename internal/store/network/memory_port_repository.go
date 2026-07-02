package network

import (
	"sync"

	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
)

type MemoryPortRepository struct {
	mu    sync.RWMutex
	ids   []string
	ports map[string]appnetwork.Port
}

func NewMemoryPortRepository() *MemoryPortRepository {
	return &MemoryPortRepository{
		ids:   []string{},
		ports: map[string]appnetwork.Port{},
	}
}

func (r *MemoryPortRepository) Create(port appnetwork.Port) appnetwork.Port {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, port.ID)
	r.ports[port.ID] = port

	return port
}

func (r *MemoryPortRepository) List() []appnetwork.Port {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ports := make([]appnetwork.Port, 0, len(r.ids))
	for _, id := range r.ids {
		ports = append(ports, r.ports[id])
	}

	return ports
}

func (r *MemoryPortRepository) Get(id string) (appnetwork.Port, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	port, ok := r.ports[id]
	if !ok {
		return appnetwork.Port{}, appnetwork.ErrPortNotFound
	}

	return port, nil
}

func (r *MemoryPortRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.ports[id]; !ok {
		return appnetwork.ErrPortNotFound
	}

	delete(r.ports, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryPortRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.ports = map[string]appnetwork.Port{}
}
