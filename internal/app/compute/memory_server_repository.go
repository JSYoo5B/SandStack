package compute

import "sync"

type MemoryServerRepository struct {
	mu      sync.RWMutex
	ids     []string
	servers map[string]Server
}

func NewMemoryServerRepository() *MemoryServerRepository {
	return &MemoryServerRepository{
		ids:     []string{},
		servers: map[string]Server{},
	}
}

func (r *MemoryServerRepository) Create(server Server) Server {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, server.ID)
	r.servers[server.ID] = server

	return server
}

func (r *MemoryServerRepository) List() []Server {
	r.mu.RLock()
	defer r.mu.RUnlock()

	servers := make([]Server, 0, len(r.ids))
	for _, id := range r.ids {
		servers = append(servers, r.servers[id])
	}

	return servers
}

func (r *MemoryServerRepository) Get(id string) (Server, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	server, ok := r.servers[id]
	if !ok {
		return Server{}, ErrServerNotFound
	}

	return server, nil
}

func (r *MemoryServerRepository) Update(server Server) (Server, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.servers[server.ID]; !ok {
		return Server{}, ErrServerNotFound
	}

	r.servers[server.ID] = server

	return server, nil
}

func (r *MemoryServerRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.servers[id]; !ok {
		return ErrServerNotFound
	}

	delete(r.servers, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryServerRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.servers = map[string]Server{}
}
