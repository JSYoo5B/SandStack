package image

type Repository interface {
	Create(image Image) Image
	List() []Image
	Get(id string) (Image, error)
	Update(image Image) (Image, error)
	Delete(id string) error
	Reset()
}
