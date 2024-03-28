package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo-app/pkg/todoItems"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todoItems.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO todo_lists (title, description) VALUES ($1, $2) RETURNING id")
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO user_lists (user_id, list_id) VALUES ($1, $2)")
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]todoItems.TodoList, error) {
	var lists []todoItems.TodoList

	getAllQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description " +
		"FROM todo_lists tl " +
		"INNER JOIN user_lists ul on tl.id = ul.list_id " +
		"WHERE ul.user_id = &1")
	err := r.db.Select(&lists, getAllQuery, userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todoItems.TodoList, error) {
	var list todoItems.TodoList

	getAllQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description " +
		"FROM todo_lists tl " +
		"INNER JOIN user_lists ul on tl.id = ul.list_id " +
		"WHERE ul.user_id = &1 AND ul.list_id = $2")
	err := r.db.Get(&list, getAllQuery, userId, listId)

	return list, err
}
