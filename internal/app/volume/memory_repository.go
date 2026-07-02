package volume

import "sync"

type MemoryRepository struct {
	mu      sync.RWMutex
	ids     []string
	volumes map[string]Volume
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		ids:     []string{},
		volumes: map[string]Volume{},
	}
}

func (r *MemoryRepository) Create(volume Volume) Volume {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, volume.ID)
	r.volumes[volume.ID] = volume

	return volume
}

func (r *MemoryRepository) List() []Volume {
	r.mu.RLock()
	defer r.mu.RUnlock()

	volumes := make([]Volume, 0, len(r.ids))
	for _, id := range r.ids {
		volumes = append(volumes, r.volumes[id])
	}

	return volumes
}

func (r *MemoryRepository) Get(id string) (Volume, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	volume, ok := r.volumes[id]
	if !ok {
		return Volume{}, ErrVolumeNotFound
	}

	return volume, nil
}

func (r *MemoryRepository) Update(volume Volume) (Volume, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.volumes[volume.ID]; !ok {
		return Volume{}, ErrVolumeNotFound
	}

	r.volumes[volume.ID] = volume

	return volume, nil
}

func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.volumes[id]; !ok {
		return ErrVolumeNotFound
	}

	delete(r.volumes, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.volumes = map[string]Volume{}
}
