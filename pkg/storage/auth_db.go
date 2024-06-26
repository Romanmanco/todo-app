package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo-app/pkg/user"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user user.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (name, username, password_hash) values ($1, $2, $3) RETURNING id")

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (user.User, error) {
	var user user.User
	query := fmt.Sprintf("SELECT id FROM users WHERE  username=$1 AND password_hash=$2")

	err := r.db.Get(&user, query, username, password)

	return user, err
}
