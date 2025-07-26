package service

import "rest_api_gin/internal/domains"

type RoleService struct {
	repo domains.RolesRepo
}

func NewRolesService(repo domains.RolesRepo) *RoleService {
	return &RoleService{repo: repo}
}

func (r *RoleService) RegisterRoleService(role *domains.Roles) error {
	if role.Name == "" {
		return nil
	}

	return r.repo.Insert(role)
}
