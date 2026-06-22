package fault

type Operation struct {
	Service string
	Name    string
}

type Rule struct {
	ID        string
	Enabled   bool
	Service   string
	Operation string
	Behavior  Behavior
	Trigger   Trigger
}

type Behavior struct {
	HTTPStatus int
	Message    string
}

type Trigger struct {
	Nth  int
	Once bool
}

type Decision struct {
	Matched    bool
	HTTPStatus int
	Message    string
}
