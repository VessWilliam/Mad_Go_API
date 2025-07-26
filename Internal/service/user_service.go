package service

import (
	"fmt"
	"rest_api_gin/internal/domains"
)

type UserService struct {
	repo domains.UserRepo
}

func NewUserService(repo domains.UserRepo) *UserService {
	return &UserService{repo: repo}
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
func (s *UserService) GetUserById(id string) (*domains.User, error) {
	if id == "" {
		return nil, fmt.Errorf("user get by id not exist: id = %q", id)
	}

	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("user get by id failed: id = %q, err: %w", id, err)
	}

	return user, nil
}
