package service

import (
	"fmt"
	"rest_api_gin/internal/domains"
)

type RoleService struct {
	repo domains.RolesRepo
}

func NewRolesService(repo domains.RolesRepo) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) RegisterRoleService(role *domains.Role) error {
	if role.Name == "" {
		return fmt.Errorf("register fails : %v", role)
	}

	return s.repo.Insert(role)
}

func (s *RoleService) GetAllRoleService() ([]*domains.Role, error) {
	roles, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("get all roles failed: %v", err)
	}
	return roles, nil
}

func (s *RoleService) GetRoleByIdService(id int) (*domains.Role, error) {
	if id == 0 {
		return nil, fmt.Errorf("role get by id not exist: id = %q", id)
	}

	role, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("role get by id failed: id = %q, err: %w", id, err)
	}

	return role, nil
}

func (s *RoleService) DeleteByIdService(id int) error {
	if id == 0 {
		return fmt.Errorf("delete role by id not exist: id = %q", id)
	}
	return s.repo.DeleteById(id)
}

func (s *RoleService) UpdateRoleService(role *domains.Role) error {
	if role == nil {
		return fmt.Errorf("update role fail: %v", role)
	}
	return s.repo.Update(role)
}
