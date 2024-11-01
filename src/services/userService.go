package services

import "github.com/Sahil2k07/go-SOLID/src/queries"

type UserHandlers interface {
	HandleSignup(name, email, password string) error
	HandleLogin(email, password string) error
	HandleUpdateUser(name string) error
}

type UserService struct {
	queries *queries.UserQueries
}

func UserServiceInit(q *queries.UserQueries) *UserService {
	return &UserService{queries: q}
}

func (us *UserService) HandleLogin(email, password string) error {
	// login

	return nil
}

func (us *UserService) HandleSignup(name, email, password string) error {
	// logic

	return nil
}

func (us *UserService) HandleUpdateUser(name string) error {
	// login

	return nil
}
