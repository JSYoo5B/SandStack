package volume

import (
	"sync"

	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
)

type MemoryAttachmentRepository struct {
	mu          sync.RWMutex
	ids         []string
	attachments map[string]appvolume.Attachment
}

func NewMemoryAttachmentRepository() *MemoryAttachmentRepository {
	return &MemoryAttachmentRepository{
		ids:         []string{},
		attachments: map[string]appvolume.Attachment{},
	}
}

func (r *MemoryAttachmentRepository) Create(
	attachment appvolume.Attachment,
) appvolume.Attachment {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, attachment.ID)
	r.attachments[attachment.ID] = attachment

	return attachment
}

func (r *MemoryAttachmentRepository) List() []appvolume.Attachment {
	r.mu.RLock()
	defer r.mu.RUnlock()

	attachments := make([]appvolume.Attachment, 0, len(r.ids))
	for _, id := range r.ids {
		attachments = append(attachments, r.attachments[id])
	}

	return attachments
}

func (r *MemoryAttachmentRepository) Get(
	id string,
) (appvolume.Attachment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	attachment, ok := r.attachments[id]
	if !ok {
		return appvolume.Attachment{}, appvolume.ErrAttachmentNotFound
	}

	return attachment, nil
}

func (r *MemoryAttachmentRepository) Update(
	attachment appvolume.Attachment,
) (appvolume.Attachment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.attachments[attachment.ID]; !ok {
		return appvolume.Attachment{}, appvolume.ErrAttachmentNotFound
	}

	r.attachments[attachment.ID] = attachment
	return attachment, nil
}

func (r *MemoryAttachmentRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.attachments[id]; !ok {
		return appvolume.ErrAttachmentNotFound
	}

	delete(r.attachments, id)
	for index, currentID := range r.ids {
		if currentID == id {
			r.ids = append(r.ids[:index], r.ids[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryAttachmentRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = []string{}
	r.attachments = map[string]appvolume.Attachment{}
}
