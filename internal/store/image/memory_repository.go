package image

import (
	"sync"

	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
)

type MemoryRepository struct {
	mu     sync.RWMutex
	ids    []string
	images map[string]appimage.Image
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		ids:    []string{},
		images: map[string]appimage.Image{},
	}
}

func (r *MemoryRepository) Create(image appimage.Image) appimage.Image {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = append(r.ids, image.ID)
	r.images[image.ID] = image

	return image
}

func (r *MemoryRepository) List() []appimage.Image {
	r.mu.RLock()
	defer r.mu.RUnlock()

	images := make([]appimage.Image, 0, len(r.ids))
	for _, id := range r.ids {
		images = append(images, r.images[id])
	}

	return images
}

func (r *MemoryRepository) Get(id string) (appimage.Image, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	image, ok := r.images[id]
	if !ok {
		return appimage.Image{}, appimage.ErrImageNotFound
	}

	return image, nil
}

func (r *MemoryRepository) Update(image appimage.Image) (appimage.Image, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.images[image.ID]; !ok {
		return appimage.Image{}, appimage.ErrImageNotFound
	}

	r.images[image.ID] = image

	return image, nil
}

func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.images[id]; !ok {
		return appimage.ErrImageNotFound
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
	r.images = map[string]appimage.Image{}
}
