package placement

import (
	"sync"

	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
)

type MemoryResourceClassRepository struct {
	mu      sync.RWMutex
	names   []string
	classes map[string]appplacement.ResourceClass
}

func NewMemoryResourceClassRepository() *MemoryResourceClassRepository {
	return &MemoryResourceClassRepository{
		names:   []string{},
		classes: map[string]appplacement.ResourceClass{},
	}
}

func (r *MemoryResourceClassRepository) Create(
	resourceClass appplacement.ResourceClass,
) appplacement.ResourceClass {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.classes[resourceClass.Name]; !ok {
		r.names = append(r.names, resourceClass.Name)
	}
	r.classes[resourceClass.Name] = resourceClass

	return resourceClass
}

func (r *MemoryResourceClassRepository) List() []appplacement.ResourceClass {
	r.mu.RLock()
	defer r.mu.RUnlock()

	classes := make([]appplacement.ResourceClass, 0, len(r.names))
	for _, name := range r.names {
		classes = append(classes, r.classes[name])
	}

	return classes
}

func (r *MemoryResourceClassRepository) Get(
	name string,
) (appplacement.ResourceClass, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	resourceClass, ok := r.classes[name]
	if !ok {
		return appplacement.ResourceClass{}, appplacement.ErrResourceClassNotFound
	}

	return resourceClass, nil
}

func (r *MemoryResourceClassRepository) Delete(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.classes[name]; !ok {
		return appplacement.ErrResourceClassNotFound
	}

	delete(r.classes, name)
	for index, currentName := range r.names {
		if currentName == name {
			r.names = append(r.names[:index], r.names[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryResourceClassRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.names = []string{}
	r.classes = map[string]appplacement.ResourceClass{}
}
