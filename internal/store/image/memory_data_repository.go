package image

import (
	"sync"

	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
)

type MemoryDataRepository struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewMemoryDataRepository() *MemoryDataRepository {
	return &MemoryDataRepository{
		data: map[string][]byte{},
	}
}

func (r *MemoryDataRepository) Put(id string, data []byte) {
	r.mu.Lock()
	defer r.mu.Unlock()

	copied := make([]byte, len(data))
	copy(copied, data)
	r.data[id] = copied
}

func (r *MemoryDataRepository) Get(id string) ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	data, ok := r.data[id]
	if !ok {
		return nil, appimage.ErrImageNotFound
	}

	copied := make([]byte, len(data))
	copy(copied, data)
	return copied, nil
}

func (r *MemoryDataRepository) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.data, id)
}

func (r *MemoryDataRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data = map[string][]byte{}
}
