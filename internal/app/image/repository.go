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
