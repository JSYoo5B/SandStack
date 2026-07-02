package compute

type ServerRepository interface {
	Create(server Server) Server
	List() []Server
	Get(id string) (Server, error)
	Update(server Server) (Server, error)
	Delete(id string) error
	Reset()
}
