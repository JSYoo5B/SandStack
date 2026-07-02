package identity

import (
	"sync"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
)

type MemoryUserRepository struct {
	mu    sync.RWMutex
	ids   []string
	users map[string]appidentity.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: map[string]appidentity.User{},
	}
}

func (r *MemoryUserRepository) Save(
	user appidentity.User,
) appidentity.User {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		r.ids = append(r.ids, user.ID)
	}
	r.users[user.ID] = user

	return user
}

func (r *MemoryUserRepository) List() []appidentity.User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]appidentity.User, 0, len(r.ids))
	for _, id := range r.ids {
		users = append(users, r.users[id])
	}

	return users
}

func (r *MemoryUserRepository) Get(
	id string,
) (appidentity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return appidentity.User{}, appidentity.ErrUserNotFound
	}

	return user, nil
}

func (r *MemoryUserRepository) FindByName(
	name string,
) (appidentity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, id := range r.ids {
		user := r.users[id]
		if user.Name == name {
			return user, nil
		}
	}

	return appidentity.User{}, appidentity.ErrUserNotFound
}

func (r *MemoryUserRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ids = nil
	r.users = map[string]appidentity.User{}
}
