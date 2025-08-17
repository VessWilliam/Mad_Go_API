package repository

import (
	"context"
	"fmt"
	"rest_api_gin/internal/domains"
	"time"

	"github.com/jmoiron/sqlx"
)

var _ domains.RolesRepo = (*RoleRepo)(nil)

type RoleRepo struct {
	DB *sqlx.DB
}

func NewRolesRepo(db *sqlx.DB) *RoleRepo {
	return &RoleRepo{DB: db}
}

func (r *RoleRepo) Insert(role *domains.Role) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into roles (name) values ($1) returning id`

	err := r.DB.QueryRowContext(ctx, query, role.Name).Scan(&role.Id)
	if err != nil {
		return fmt.Errorf("insert role fail %v", err)
	}

	return nil
}

func (r *RoleRepo) GetAll() ([]*domains.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var roles []*domains.Role
	query := `select * from roles`

	err := r.DB.SelectContext(ctx, &roles, query)
	if err != nil {
		return nil, fmt.Errorf("get all roles %v", err)
	}

	return roles, nil
}

func (r *RoleRepo) GetById(id int) (*domains.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var role domains.Role
	query := `select * from roles where id = $1`

	err := r.DB.GetContext(ctx, &role, query, id)
	if err != nil {
		return nil, fmt.Errorf("get role by id fail %v", err)
	}

	return &role, nil
}

func (r *RoleRepo) DeleteById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `delete from roles where id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete role by id fail %v", err)
	}

	return nil
}

func (r *RoleRepo) Update(role *domains.Role) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update roles set name = $1 where id = $2`
	_, err := r.DB.ExecContext(ctx, query, role.Name, role.Id)
	if err != nil {
		return fmt.Errorf("update role fail :%v", err)
	}
	return nil
}

func (r *RoleRepo) GetRoleByEmail(email string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
    SELECT r.name
	FROM user_roles ur
	JOIN users u ON ur.user_id = u.id
	JOIN roles r ON ur.role_id = r.id
	WHERE u.email = $1;
   `

	var roles []string
	if err := r.DB.SelectContext(ctx, &roles, query, email); err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		roles = []string{"Unassign"}
	}

	return roles, nil

}
