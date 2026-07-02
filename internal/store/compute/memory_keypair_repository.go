package compute

import (
	"sync"

	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
)

type MemoryKeyPairRepository struct {
	mu       sync.RWMutex
	names    []string
	keyPairs map[string]appcompute.KeyPair
}

func NewMemoryKeyPairRepository() *MemoryKeyPairRepository {
	return &MemoryKeyPairRepository{
		names:    []string{},
		keyPairs: map[string]appcompute.KeyPair{},
	}
}

func (r *MemoryKeyPairRepository) Create(
	keyPair appcompute.KeyPair,
) appcompute.KeyPair {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.keyPairs[keyPair.Name]; !ok {
		r.names = append(r.names, keyPair.Name)
	}
	r.keyPairs[keyPair.Name] = keyPair

	return keyPair
}

func (r *MemoryKeyPairRepository) List() []appcompute.KeyPair {
	r.mu.RLock()
	defer r.mu.RUnlock()

	keyPairs := make([]appcompute.KeyPair, 0, len(r.names))
	for _, name := range r.names {
		keyPairs = append(keyPairs, r.keyPairs[name])
	}

	return keyPairs
}

func (r *MemoryKeyPairRepository) Get(name string) (appcompute.KeyPair, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	keyPair, ok := r.keyPairs[name]
	if !ok {
		return appcompute.KeyPair{}, appcompute.ErrKeyPairNotFound
	}

	return keyPair, nil
}

func (r *MemoryKeyPairRepository) Delete(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.keyPairs[name]; !ok {
		return appcompute.ErrKeyPairNotFound
	}

	delete(r.keyPairs, name)
	for index, currentName := range r.names {
		if currentName == name {
			r.names = append(r.names[:index], r.names[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryKeyPairRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.names = []string{}
	r.keyPairs = map[string]appcompute.KeyPair{}
}
