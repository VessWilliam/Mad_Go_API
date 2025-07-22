package repository

import (
	"context"
	utils "rest_api_gin/internal/Utils"
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

	return utils.NameGetContext(ctx, r.DB, query, &user.Id, user)
}

// Get implements domains.UserRepo.
func (r *UserRepo) Get(id int) (*domains.User, error) {
	panic("unimplemented")
}
