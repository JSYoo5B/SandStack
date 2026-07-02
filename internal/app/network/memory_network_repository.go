package network

import "sync"

type MemoryNetworkRepository struct {
	mu       sync.RWMutex
	ids      []string
	networks map[string]Network
}

func NewMemoryNetworkRepository() *MemoryNetworkRepository {
	return &MemoryNetworkRepository{
		ids:      []string{},
		networks: map[string]Network{},
	}
}

func (r *MemoryNetworkRepository) Create(network Network) Network {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, network.ID)
	r.networks[network.ID] = network

	return network
}

func (r *MemoryNetworkRepository) List() []Network {
	r.mu.RLock()
	defer r.mu.RUnlock()

	networks := make([]Network, 0, len(r.ids))
	for _, id := range r.ids {
		networks = append(networks, r.networks[id])
	}

	return networks
}

func (r *MemoryNetworkRepository) Get(id string) (Network, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	network, ok := r.networks[id]
	if !ok {
		return Network{}, ErrNetworkNotFound
	}

	return network, nil
}

func (r *MemoryNetworkRepository) Update(network Network) (Network, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.networks[network.ID]; !ok {
		return Network{}, ErrNetworkNotFound
	}

	r.networks[network.ID] = network

	return network, nil
}

func (r *MemoryNetworkRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.networks[id]; !ok {
		return ErrNetworkNotFound
	}

	delete(r.networks, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryNetworkRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.networks = map[string]Network{}
}
