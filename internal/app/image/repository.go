package image

type Repository interface {
	Create(image Image) Image
	List() []Image
	Get(id string) (Image, error)
	Update(image Image) (Image, error)
	Delete(id string) error
	Reset()
}

type DataRepository interface {
	Put(id string, data []byte)
	Get(id string) ([]byte, error)
	Delete(id string)
	Reset()
}

type MemberRepository interface {
	Create(member Member) Member
	List(imageID string) []Member
	Get(imageID string, memberID string) (Member, error)
	Update(member Member) (Member, error)
	Delete(imageID string, memberID string) error
	Reset()
}

type TaskRepository interface {
	Create(task Task) Task
	List() []Task
	Get(id string) (Task, error)
	Reset()
}
