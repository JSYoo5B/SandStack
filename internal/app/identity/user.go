package identity

func (s Service) Users() []User {
	return s.repositories.Users.List()
}

func (s Service) defaultUser() User {
	return User{
		DefaultProjectID: s.config.ProjectID,
		Description:      "Default SandStack user",
		DomainID:         "default",
		Enabled:          true,
		ID:               s.config.UserID,
		Name:             s.config.Username,
		Password:         s.config.Password,
	}
}

func (s Service) UserByID(id string) (User, bool) {
	user, err := s.repositories.Users.Get(id)
	if err != nil {
		return User{}, false
	}

	return user, true
}
