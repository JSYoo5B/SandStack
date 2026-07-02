package network

import (
	"sync"

	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
)

type MemoryRouterInterfaceRepository struct {
	mu         sync.RWMutex
	ids        []string
	interfaces map[string]appnetwork.RouterInterface
}

func NewMemoryRouterInterfaceRepository() *MemoryRouterInterfaceRepository {
	return &MemoryRouterInterfaceRepository{
		ids:        []string{},
		interfaces: map[string]appnetwork.RouterInterface{},
	}
}

func (r *MemoryRouterInterfaceRepository) Create(
	routerInterface appnetwork.RouterInterface,
) appnetwork.RouterInterface {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, routerInterface.ID)
	r.interfaces[routerInterface.ID] = routerInterface

	return routerInterface
}

func (r *MemoryRouterInterfaceRepository) Find(
	routerID string,
	request appnetwork.RouterInterfaceRequest,
) (appnetwork.RouterInterface, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, id := range r.ids {
		routerInterface := r.interfaces[id]
		if routerInterface.RouterID != routerID {
			continue
		}
		if request.PortID != "" && routerInterface.PortID == request.PortID {
			return routerInterface, nil
		}
		if request.SubnetID != "" && routerInterface.SubnetID == request.SubnetID {
			return routerInterface, nil
		}
	}

	return appnetwork.RouterInterface{}, appnetwork.ErrRouterInterfaceNotFound
}

func (r *MemoryRouterInterfaceRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.interfaces[id]; !ok {
		return appnetwork.ErrRouterInterfaceNotFound
	}

	delete(r.interfaces, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryRouterInterfaceRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.interfaces = map[string]appnetwork.RouterInterface{}
}
