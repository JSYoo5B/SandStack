package network

import (
	"sync"

	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
)

type MemoryNetworkRepository struct {
	mu       sync.RWMutex
	ids      []string
	networks map[string]appnetwork.Network
}

func NewMemoryNetworkRepository() *MemoryNetworkRepository {
	return &MemoryNetworkRepository{
		ids:      []string{},
		networks: map[string]appnetwork.Network{},
	}
}

func (r *MemoryNetworkRepository) Create(network appnetwork.Network) appnetwork.Network {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, network.ID)
	r.networks[network.ID] = network

	return network
}

func (r *MemoryNetworkRepository) List() []appnetwork.Network {
	r.mu.RLock()
	defer r.mu.RUnlock()

	networks := make([]appnetwork.Network, 0, len(r.ids))
	for _, id := range r.ids {
		networks = append(networks, r.networks[id])
	}

	return networks
}

func (r *MemoryNetworkRepository) Get(id string) (appnetwork.Network, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	network, ok := r.networks[id]
	if !ok {
		return appnetwork.Network{}, appnetwork.ErrNetworkNotFound
	}

	return network, nil
}

func (r *MemoryNetworkRepository) Update(network appnetwork.Network) (appnetwork.Network, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.networks[network.ID]; !ok {
		return appnetwork.Network{}, appnetwork.ErrNetworkNotFound
	}

	r.networks[network.ID] = network

	return network, nil
}

func (r *MemoryNetworkRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.networks[id]; !ok {
		return appnetwork.ErrNetworkNotFound
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
	r.networks = map[string]appnetwork.Network{}
}
