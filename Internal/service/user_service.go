package service

import "rest_api_gin/internal/domains"

type UserService struct {
	repo domains.UserRepo
}

func NewUserService(repo domains.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *domains.User) error {

	if user.Email == "" || user.Password == "" {
		return nil
	}

	return s.repo.Insert(user)
}
