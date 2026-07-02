package volume

type Repository interface {
	Create(volume Volume) Volume
	List() []Volume
	Get(id string) (Volume, error)
	Update(volume Volume) (Volume, error)
	Delete(id string) error
	Reset()
}
