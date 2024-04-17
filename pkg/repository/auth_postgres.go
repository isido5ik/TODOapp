package repository

import (
	"fmt"

	"github.com/isido5ik/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id;", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

// to delete
// func (r *AuthPostgres) GetUsers() ([]todo.User, error) {
// 	query := fmt.Sprintf("SELECT id, name, username FROM users")
// 	rows, err := r.db.Queryx(query)
// 	if err != nil {
// 		logrus.Errorf("Ошибка при выполнении запроса GetUsers: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []todo.User
// 	for rows.Next() {
// 		var user todo.User
// 		if err := rows.StructScan(&user); err != nil {
// 			logrus.Errorf("Ошибка при сканировании строки в GetUsers: %v", err)
// 			continue
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }
