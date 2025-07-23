package repository

import (
	"context"
	"fmt"
	"rest_api_gin/internal/domains"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Insert(user *domains.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	   Insert into users (email, password, name) 
	   values (:email, :password, :name) returning id
	`

	row, err := r.DB.NamedQueryContext(ctx, query, user)
	if err != nil {
		return err
	}
	defer row.Close()

	if row.Next() {
		return row.Scan(&user.Id)
	}

	return fmt.Errorf("Insert fail: no Id returned")
}

func (r *UserRepo) GetAll() ([]*domains.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var users []*domains.User
	query := `select * from users`

	err := r.DB.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}

	return users, err
}
