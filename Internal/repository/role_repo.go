package repository

import (
	"github.com/jmoiron/sqlx"
)

type RolesRepo struct {
	DB *sqlx.DB
}

func NewRolesRepo(db *sqlx.DB) *RolesRepo {
	return &RolesRepo{DB: db}
}
