package storage

import (
	"github.com/jmoiron/sqlx"
	"todo-app/pkg/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GetUser(username, password string) (user.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
