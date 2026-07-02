package volume

import (
	"sync"

	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
)

type MemoryBackupRepository struct {
	mu      sync.RWMutex
	ids     []string
	backups map[string]appvolume.Backup
}

func NewMemoryBackupRepository() *MemoryBackupRepository {
	return &MemoryBackupRepository{
		ids:     []string{},
		backups: map[string]appvolume.Backup{},
	}
}

func (r *MemoryBackupRepository) Create(backup appvolume.Backup) appvolume.Backup {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, backup.ID)
	r.backups[backup.ID] = backup

	return backup
}

func (r *MemoryBackupRepository) List() []appvolume.Backup {
	r.mu.RLock()
	defer r.mu.RUnlock()

	backups := make([]appvolume.Backup, 0, len(r.ids))
	for _, id := range r.ids {
		backups = append(backups, r.backups[id])
	}

	return backups
}

func (r *MemoryBackupRepository) Get(id string) (appvolume.Backup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	backup, ok := r.backups[id]
	if !ok {
		return appvolume.Backup{}, appvolume.ErrBackupNotFound
	}

	return backup, nil
}

func (r *MemoryBackupRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.backups[id]; !ok {
		return appvolume.ErrBackupNotFound
	}

	delete(r.backups, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryBackupRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.backups = map[string]appvolume.Backup{}
}
