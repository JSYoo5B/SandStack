package network

type NetworkRepository interface {
	Create(network Network) Network
	List() []Network
	Get(id string) (Network, error)
	Update(network Network) (Network, error)
	Delete(id string) error
	Reset()
}

type SubnetRepository interface {
	Create(subnet Subnet) Subnet
	List() []Subnet
	Get(id string) (Subnet, error)
	Delete(id string) error
	Reset()
}
