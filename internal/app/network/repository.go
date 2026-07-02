package network

type NetworkRepository interface {
	Create(network Network) Network
	List() []Network
	Get(id string) (Network, error)
	Update(network Network) (Network, error)
	Delete(id string) error
	Reset()
}
