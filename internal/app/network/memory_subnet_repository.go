package network

import "sync"

type MemorySubnetRepository struct {
	mu      sync.RWMutex
	ids     []string
	subnets map[string]Subnet
}

func NewMemorySubnetRepository() *MemorySubnetRepository {
	return &MemorySubnetRepository{
		ids:     []string{},
		subnets: map[string]Subnet{},
	}
}

func (r *MemorySubnetRepository) Create(subnet Subnet) Subnet {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, subnet.ID)
	r.subnets[subnet.ID] = subnet

	return subnet
}

func (r *MemorySubnetRepository) List() []Subnet {
	r.mu.RLock()
	defer r.mu.RUnlock()

	subnets := make([]Subnet, 0, len(r.ids))
	for _, id := range r.ids {
		subnets = append(subnets, r.subnets[id])
	}

	return subnets
}

func (r *MemorySubnetRepository) Get(id string) (Subnet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	subnet, ok := r.subnets[id]
	if !ok {
		return Subnet{}, ErrSubnetNotFound
	}

	return subnet, nil
}

func (r *MemorySubnetRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.subnets[id]; !ok {
		return ErrSubnetNotFound
	}

	delete(r.subnets, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemorySubnetRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.subnets = map[string]Subnet{}
}
