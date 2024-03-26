package service

import (
	"todo-app/pkg/storage"
	"todo-app/pkg/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *storage.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
