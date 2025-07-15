package database

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "Insert into users (email, password, name) values ($1,$2,$3) returning id"

	err := m.DB.QueryRowContext(ctx, query, user.Email, user.Password, user.Name).Scan(&user.Id)
	if err != nil {
		log.Println("Insert User Error:", err)
	}

	return err

}

func (m *UserModel) Get(id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	return m.getUser(query, id)
}

func (m *UserModel) getUser(query string, args ...interface{}) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
