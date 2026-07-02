package volume

import (
	"sync"

	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
)

type MemoryTransferRepository struct {
	mu        sync.RWMutex
	ids       []string
	transfers map[string]appvolume.Transfer
}

func NewMemoryTransferRepository() *MemoryTransferRepository {
	return &MemoryTransferRepository{
		ids:       []string{},
		transfers: map[string]appvolume.Transfer{},
	}
}

func (r *MemoryTransferRepository) Create(
	transfer appvolume.Transfer,
) appvolume.Transfer {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, transfer.ID)
	r.transfers[transfer.ID] = transfer

	return transfer
}

func (r *MemoryTransferRepository) List() []appvolume.Transfer {
	r.mu.RLock()
	defer r.mu.RUnlock()

	transfers := make([]appvolume.Transfer, 0, len(r.ids))
	for _, id := range r.ids {
		transfers = append(transfers, r.transfers[id])
	}

	return transfers
}

func (r *MemoryTransferRepository) Get(id string) (appvolume.Transfer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	transfer, ok := r.transfers[id]
	if !ok {
		return appvolume.Transfer{}, appvolume.ErrTransferNotFound
	}

	return transfer, nil
}

func (r *MemoryTransferRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.transfers[id]; !ok {
		return appvolume.ErrTransferNotFound
	}

	delete(r.transfers, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryTransferRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.transfers = map[string]appvolume.Transfer{}
}
