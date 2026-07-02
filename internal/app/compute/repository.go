package compute

type ServerRepository interface {
	Create(server Server) Server
	List() []Server
	Get(id string) (Server, error)
	Update(server Server) (Server, error)
	Delete(id string) error
	Reset()
}

type KeyPairRepository interface {
	Create(keyPair KeyPair) KeyPair
	List() []KeyPair
	Get(name string) (KeyPair, error)
	Delete(name string) error
	Reset()
}
