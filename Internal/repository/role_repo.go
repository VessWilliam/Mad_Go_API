package repository

import (
	"context"
	"fmt"
	"rest_api_gin/internal/domains"
	"time"

	"github.com/jmoiron/sqlx"
)

type RolesRepo struct {
	DB *sqlx.DB
}

func NewRolesRepo(db *sqlx.DB) *RolesRepo {
	return &RolesRepo{DB: db}
}

func (r *RolesRepo) Insert(role *domains.Roles) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into roles (name) values ($1) returning id`

	err := r.DB.QueryRowContext(ctx, query, role.Name).Scan(&role.Id)
	if err != nil {
		return fmt.Errorf("insert role fail %v", err)
	}

	return nil
}
