package volume

import (
	"sync"

	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
)

type MemoryRepository struct {
	mu      sync.RWMutex
	ids     []string
	volumes map[string]appvolume.Volume
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		ids:     []string{},
		volumes: map[string]appvolume.Volume{},
	}
}

func (r *MemoryRepository) Create(volume appvolume.Volume) appvolume.Volume {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, volume.ID)
	r.volumes[volume.ID] = volume

	return volume
}

func (r *MemoryRepository) List() []appvolume.Volume {
	r.mu.RLock()
	defer r.mu.RUnlock()

	volumes := make([]appvolume.Volume, 0, len(r.ids))
	for _, id := range r.ids {
		volumes = append(volumes, r.volumes[id])
	}

	return volumes
}

func (r *MemoryRepository) Get(id string) (appvolume.Volume, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	volume, ok := r.volumes[id]
	if !ok {
		return appvolume.Volume{}, appvolume.ErrVolumeNotFound
	}

	return volume, nil
}

func (r *MemoryRepository) Update(
	volume appvolume.Volume,
) (appvolume.Volume, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.volumes[volume.ID]; !ok {
		return appvolume.Volume{}, appvolume.ErrVolumeNotFound
	}

	r.volumes[volume.ID] = volume

	return volume, nil
}

func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.volumes[id]; !ok {
		return appvolume.ErrVolumeNotFound
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
	r.volumes = map[string]appvolume.Volume{}
}
