package image

import "sync"

type MemoryRepository struct {
	mu     sync.RWMutex
	ids    []string
	images map[string]Image
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		ids:    []string{},
		images: map[string]Image{},
	}
}

func (r *MemoryRepository) Create(image Image) Image {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, image.ID)
	r.images[image.ID] = image

	return image
}

func (r *MemoryRepository) List() []Image {
	r.mu.RLock()
	defer r.mu.RUnlock()

	images := make([]Image, 0, len(r.ids))
	for _, id := range r.ids {
		images = append(images, r.images[id])
	}

	return images
}

func (r *MemoryRepository) Get(id string) (Image, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	image, ok := r.images[id]
	if !ok {
		return Image{}, ErrImageNotFound
	}

	return image, nil
}

func (r *MemoryRepository) Update(image Image) (Image, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.images[image.ID]; !ok {
		return Image{}, ErrImageNotFound
	}

	r.images[image.ID] = image

	return image, nil
}

func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.images[id]; !ok {
		return ErrImageNotFound
	}

	delete(r.images, id)
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
	r.images = map[string]Image{}
}
