package queries

import "database/sql"

type UserMethods interface {
	CreateUser(name, email, password string) error
	UpdateUser(name, email, password string) error
}

type UserQueries struct {
	db *sql.DB
}

func UserQueriesInit(db *sql.DB) *UserQueries {
	return &UserQueries{db: db}
}

func (uq *UserQueries) CreateUser(name, email, password string) error {
	// logic

	return nil
}

func (uq *UserQueries) UpdateUser(name, email, password string) error {
	// logic

	return nil
}
