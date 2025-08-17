package repository

import (
	"context"
	"fmt"
	"rest_api_gin/internal/domains"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var _ domains.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Insert(user *domains.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	query := `
	   Insert into users (email, password, name) 
	   values ($1, $2, $3) returning id
	`

	err = r.DB.QueryRowContext(ctx, query, user.Email, string(hashPass), user.Name).Scan(&user.Id)
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

func (r *UserRepo) GetById(id int) (*domains.User, error) {
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

func (r *UserRepo) GetRolesByUserId(userId int) ([]domains.Role, error) {
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

func (r *UserRepo) AssignRolesToRoles(userId int, roleIds []int) error {

	tx, err := r.DB.Beginx()
	if err != nil {
		return fmt.Errorf("assign row failed : %v ", err)
	}

	defer tx.Rollback()

	_, err = tx.Exec(`delete from user_roles where user_id = $1`, userId)
	if err != nil {
		return fmt.Errorf("delete user_id in user_roles : %v ", err)
	}

	for _, role_id := range roleIds {
		_, err := tx.Exec(
			`insert into user_roles (user_id, role_id) values ($1 , $2)`,
			userId, role_id,
		)

		if err != nil {
			return fmt.Errorf("insert user_roles error : %v", err)
		}
	}

	return tx.Commit()
}

func (r *UserRepo) Update(user *domains.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	  update users SET name = $1, email = $2 WHERE id = $3;
	`
	_, err := r.DB.ExecContext(ctx, query, user.Name, user.Email, user.Id)
	if err != nil {
		return fmt.Errorf("update users error : %v", err)
	}
	return nil
}

func (r *UserRepo) GetByEmail(email string) (*domains.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	    select email, password, name
		from users where email = $1
	`

	user := &domains.User{}
	err := r.DB.QueryRowContext(ctx, query, email).Scan(&user.Email, &user.Password, &user.Name)

	if err != nil {
		return nil, fmt.Errorf("user not found %v", err)
	}

	return user, nil
}
