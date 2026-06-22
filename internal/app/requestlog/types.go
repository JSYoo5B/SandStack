package requestlog

type Record struct {
	ID     string
	Method string
	Path   string
	Status int
}
