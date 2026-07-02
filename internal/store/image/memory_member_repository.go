package image

import (
	"sync"

	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
)

type MemoryMemberRepository struct {
	mu      sync.RWMutex
	keys    []memberKey
	members map[memberKey]appimage.Member
}

type memberKey struct {
	imageID  string
	memberID string
}

func NewMemoryMemberRepository() *MemoryMemberRepository {
	return &MemoryMemberRepository{
		keys:    []memberKey{},
		members: map[memberKey]appimage.Member{},
	}
}

func (r *MemoryMemberRepository) Create(member appimage.Member) appimage.Member {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := memberKey{imageID: member.ImageID, memberID: member.MemberID}
	if _, ok := r.members[key]; !ok {
		r.keys = append(r.keys, key)
	}
	r.members[key] = member

	return member
}

func (r *MemoryMemberRepository) List(imageID string) []appimage.Member {
	r.mu.RLock()
	defer r.mu.RUnlock()

	members := []appimage.Member{}
	for _, key := range r.keys {
		if key.imageID == imageID {
			members = append(members, r.members[key])
		}
	}

	return members
}

func (r *MemoryMemberRepository) Get(
	imageID string,
	memberID string,
) (appimage.Member, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	member, ok := r.members[memberKey{imageID: imageID, memberID: memberID}]
	if !ok {
		return appimage.Member{}, appimage.ErrImageNotFound
	}

	return member, nil
}

func (r *MemoryMemberRepository) Update(
	member appimage.Member,
) (appimage.Member, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := memberKey{imageID: member.ImageID, memberID: member.MemberID}
	if _, ok := r.members[key]; !ok {
		return appimage.Member{}, appimage.ErrImageNotFound
	}
	r.members[key] = member

	return member, nil
}

func (r *MemoryMemberRepository) Delete(imageID string, memberID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := memberKey{imageID: imageID, memberID: memberID}
	if _, ok := r.members[key]; !ok {
		return appimage.ErrImageNotFound
	}

	delete(r.members, key)
	for index, currentKey := range r.keys {
		if currentKey == key {
			r.keys = append(r.keys[:index], r.keys[index+1:]...)
			break
		}
	}

	return nil
}

func (r *MemoryMemberRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.keys = []memberKey{}
	r.members = map[memberKey]appimage.Member{}
}
