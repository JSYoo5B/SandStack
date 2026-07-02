package volume

import (
	"sync"

	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
)

type MemorySnapshotRepository struct {
	mu        sync.RWMutex
	ids       []string
	snapshots map[string]appvolume.Snapshot
}

func NewMemorySnapshotRepository() *MemorySnapshotRepository {
	return &MemorySnapshotRepository{
		ids:       []string{},
		snapshots: map[string]appvolume.Snapshot{},
	}
}

func (r *MemorySnapshotRepository) Create(
	snapshot appvolume.Snapshot,
) appvolume.Snapshot {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, snapshot.ID)
	r.snapshots[snapshot.ID] = snapshot

	return snapshot
}

func (r *MemorySnapshotRepository) List() []appvolume.Snapshot {
	r.mu.RLock()
	defer r.mu.RUnlock()

	snapshots := make([]appvolume.Snapshot, 0, len(r.ids))
	for _, id := range r.ids {
		snapshots = append(snapshots, r.snapshots[id])
	}

	return snapshots
}

func (r *MemorySnapshotRepository) Get(id string) (appvolume.Snapshot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	snapshot, ok := r.snapshots[id]
	if !ok {
		return appvolume.Snapshot{}, appvolume.ErrSnapshotNotFound
	}

	return snapshot, nil
}

func (r *MemorySnapshotRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.snapshots[id]; !ok {
		return appvolume.ErrSnapshotNotFound
	}

	delete(r.snapshots, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemorySnapshotRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.snapshots = map[string]appvolume.Snapshot{}
}
