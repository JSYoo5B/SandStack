package identity

type User struct {
	DefaultProjectID string
	Description      string
	DomainID         string
	Enabled          bool
	ID               string
	Name             string
}

func (s Service) Users() []User {
	return []User{s.User()}
}

func (s Service) User() User {
	return User{
		DefaultProjectID: s.config.ProjectID,
		Description:      "Default SandStack user",
		DomainID:         "default",
		Enabled:          true,
		ID:               s.config.UserID,
		Name:             s.config.Username,
	}
}

func (s Service) UserByID(id string) (User, bool) {
	user := s.User()
	if user.ID != id {
		return User{}, false
	}

	return user, true
}
