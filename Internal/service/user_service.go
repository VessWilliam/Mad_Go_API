package service

import (
	"fmt"
	"rest_api_gin/internal/domains"
)

type UserService struct {
	repo     domains.UserRepo
	roleRepo domains.RolesRepo
}

func NewUserService(repo domains.UserRepo, roleRepo domains.RolesRepo) *UserService {
	return &UserService{
		repo:     repo,
		roleRepo: roleRepo,
	}
}

func (s *UserService) RegisterUser(user *domains.User) error {

	if user.Email == "" || user.Password == "" {
		return fmt.Errorf("register fails : %v / %v",
			user.Email, user.Password)
	}

	return s.repo.Insert(user)
}

func (s *UserService) GetAllUser() ([]*domains.User, error) {

	users, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("get all users failed: %v", err)
	}
	return users, nil

}
func (s *UserService) GetUserById(id int) (*domains.User, error) {
	if id == 0 {
		return nil, fmt.Errorf("user get by id not exist: id = %q", id)
	}

	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("user get by id failed: id = %q, err: %w", id, err)
	}

	roles, err := s.repo.GetRolesByUserId(id)
	if err != nil {
		return nil, fmt.Errorf("user get role failed: %v", err)
	}

	user.Role = roles
	return user, nil
}

func (s *UserService) AssignRolesToUser(userId int, roleIds []int) error {
	if userId == 0 || len(roleIds) == 0 {
		return fmt.Errorf("")
	}

	user, err := s.repo.GetById(userId)
	if err != nil {
		return fmt.Errorf("user not found: id = %q, err: %w", userId, err)
	}
	for _, roleId := range roleIds {
		_, err := s.roleRepo.GetById(roleId)
		if err != nil {
			return fmt.Errorf("invalid roleId %q: %w", roleId, err)
		}
	}

	err = s.repo.AssignRolesToRoles(user.Id, roleIds)
	if err != nil {
		return fmt.Errorf("failed to assign roles: %w", err)
	}

	return nil
}
