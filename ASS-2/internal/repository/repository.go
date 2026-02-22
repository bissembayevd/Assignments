package repository

import (
	_postgres "ASS-2/internal/repository/postgres"
	"ASS-2/internal/repository/postgres/users"
	"ASS-2/pkg/modules"
)

type UserRepository interface {
	GetUsers() ([]modules.User, error)
	GetUserByID(id int) (*modules.User, error)
	CreateUser(input modules.CreateUserInput) (int, error)
	UpdateUser(id int, input modules.UpdateUserInput) error
	DeleteUser(id int) (int64, error)
}

type Repositories struct {
	UserRepository
}

func NewRepositories(db *_postgres.Dialect) *Repositories {
	return &Repositories{
		UserRepository: users.NewUserRepository(db),
	}
}
