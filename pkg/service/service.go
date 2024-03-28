package service

import (
	"todo-app/pkg/storage"
	"todo-app/pkg/todoItems"
	"todo-app/pkg/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todoItems.TodoList) (int, error)
	GetAll(userId int) ([]todoItems.TodoList, error)
	GetById(userId, listId int) (todoItems.TodoList, error)
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
		TodoList:      NewTodoListService(repo),
	}
}
