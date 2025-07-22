package database

import (
	"context"
	"database/sql"
	"log"
	utils "rest_api_gin/internal/Utils"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserModel struct {
	DB *sqlx.DB
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

	query := `Insert into users (email, password, name) 
	 values (:email,:password,:name)
	 returning id`

	result := utils.NameGetContext(ctx, m.DB, query, &user.Id, user)
	if result != nil {
		log.Println("Insert User Error:", result)
	}
	return result
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
