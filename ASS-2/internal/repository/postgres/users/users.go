package users

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_postgres "ASS-2/internal/repository/postgres"
	"ASS-2/pkg/modules"
)

var ErrNotFound = errors.New("user not found")

type Repository struct {
	db               *_postgres.Dialect
	executionTimeout time.Duration
}

func NewUserRepository(db *_postgres.Dialect) *Repository {
	return &Repository{
		db:               db,
		executionTimeout: time.Second * 5,
	}
}

func (r *Repository) GetUsers() ([]modules.User, error) {
	var users []modules.User
	err := r.db.DB.Select(&users, "SELECT id, name, email, age, created_at FROM users ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("GetUsers: %w", err)
	}
	return users, nil
}

func (r *Repository) GetUserByID(id int) (*modules.User, error) {
	var user modules.User
	err := r.db.DB.Get(&user, "SELECT id, name, email, age, created_at FROM users WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}
	return &user, nil
}

func (r *Repository) CreateUser(input modules.CreateUserInput) (int, error) {
	var id int
	err := r.db.DB.QueryRow(
		"INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id",
		input.Name, input.Email, input.Age,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}
	return id, nil
}

func (r *Repository) UpdateUser(id int, input modules.UpdateUserInput) error {
	result, err := r.db.DB.Exec(
		"UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4",
		input.Name, input.Email, input.Age, id,
	)
	if err != nil {
		return fmt.Errorf("UpdateUser: %w", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *Repository) DeleteUser(id int) (int64, error) {
	result, err := r.db.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return 0, fmt.Errorf("DeleteUser: %w", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return 0, ErrNotFound
	}
	return rows, nil
}
