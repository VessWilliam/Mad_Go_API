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
	   values ($1, $2, $3) returning id
	`

	err := r.DB.QueryRowContext(ctx, query, user.Email, user.Password, user.Name).Scan(&user.Id)
	if err != nil {
		return fmt.Errorf("insert failed %v", err)
	}

	return nil
}

func (r *UserRepo) GetAll() ([]*domains.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var users []*domains.User
	query := `select * from users`

	err := r.DB.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, fmt.Errorf("get all user %v", err)
	}

	return users, nil
}

func (r *UserRepo) GetById(id string) (*domains.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user domains.User
	query := `select * from users where id = $1`

	err := r.DB.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, fmt.Errorf("get user by id %v", err)
	}

	return &user, nil
}

func (r *UserRepo) GetRolesByUserId(userId string) ([]domains.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select r.id, r.name 
	from roles r inner join user_roles ur 
	on r.id = ur.role_id 
	where user_id = $1`

	var roles []domains.Role
	if err := r.DB.SelectContext(ctx, &roles, query, userId); err != nil {
		return nil, fmt.Errorf("get roles for user %v failed: %w", userId, err)
	}

	return roles, nil
}
