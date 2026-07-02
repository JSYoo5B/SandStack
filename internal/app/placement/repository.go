package placement

type ResourceProviderRepository interface {
	Create(provider ResourceProvider) ResourceProvider
	List() []ResourceProvider
	Get(uuid string) (ResourceProvider, error)
	Delete(uuid string) error
	Reset()
}
