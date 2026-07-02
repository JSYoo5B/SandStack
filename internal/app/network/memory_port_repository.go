package network

import "sync"

type MemoryPortRepository struct {
	mu    sync.RWMutex
	ids   []string
	ports map[string]Port
}

func NewMemoryPortRepository() *MemoryPortRepository {
	return &MemoryPortRepository{
		ids:   []string{},
		ports: map[string]Port{},
	}
}

func (r *MemoryPortRepository) Create(port Port) Port {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, port.ID)
	r.ports[port.ID] = port

	return port
}

func (r *MemoryPortRepository) List() []Port {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ports := make([]Port, 0, len(r.ids))
	for _, id := range r.ids {
		ports = append(ports, r.ports[id])
	}

	return ports
}

func (r *MemoryPortRepository) Get(id string) (Port, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	port, ok := r.ports[id]
	if !ok {
		return Port{}, ErrPortNotFound
	}

	return port, nil
}

func (r *MemoryPortRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.ports[id]; !ok {
		return ErrPortNotFound
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
	r.ports = map[string]Port{}
}
