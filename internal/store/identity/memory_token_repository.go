package identity

import (
	"sync"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
)

type MemoryTokenRepository struct {
	mu     sync.RWMutex
	tokens map[string]appidentity.IssuedToken
}

func NewMemoryTokenRepository() *MemoryTokenRepository {
	return &MemoryTokenRepository{
		tokens: map[string]appidentity.IssuedToken{},
	}
}

func (r *MemoryTokenRepository) Save(
	token appidentity.IssuedToken,
) appidentity.IssuedToken {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tokens[token.ID] = token

	return token
}

func (r *MemoryTokenRepository) Get(
	id string,
) (appidentity.IssuedToken, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	token, ok := r.tokens[id]
	if !ok {
		return appidentity.IssuedToken{}, appidentity.ErrTokenNotFound
	}

	return token, nil
}

func (r *MemoryTokenRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tokens[id]; !ok {
		return appidentity.ErrTokenNotFound
	}
	delete(r.tokens, id)

	return nil
}

func (r *MemoryTokenRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tokens = map[string]appidentity.IssuedToken{}
}
