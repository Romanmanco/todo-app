package service

import (
	"todo-app/pkg/storage"
	"todo-app/pkg/todoItems"
)

type TodoListService struct {
	repo storage.TodoList
}

func NewTodoListService(repo storage.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(userId int, list todoItems.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
