package placement

import (
	"sync"

	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
)

type MemoryTraitCatalogRepository struct {
	mu     sync.RWMutex
	names  []string
	traits map[string]string
}

func NewMemoryTraitCatalogRepository() *MemoryTraitCatalogRepository {
	return &MemoryTraitCatalogRepository{
		names:  []string{},
		traits: map[string]string{},
	}
}

func (r *MemoryTraitCatalogRepository) Create(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.traits[name]; !ok {
		r.names = append(r.names, name)
	}
	r.traits[name] = name
}

func (r *MemoryTraitCatalogRepository) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	traits := make([]string, 0, len(r.names))
	for _, name := range r.names {
		traits = append(traits, r.traits[name])
	}

	return traits
}

func (r *MemoryTraitCatalogRepository) Get(name string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	trait, ok := r.traits[name]
	if !ok {
		return "", appplacement.ErrTraitNotFound
	}

	return trait, nil
}

func (r *MemoryTraitCatalogRepository) Delete(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.traits[name]; !ok {
		return appplacement.ErrTraitNotFound
	}

	delete(r.traits, name)
	for index, currentName := range r.names {
		if currentName == name {
			r.names = append(r.names[:index], r.names[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryTraitCatalogRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.names = []string{}
	r.traits = map[string]string{}
}
