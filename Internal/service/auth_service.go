package service

import (
	"fmt"
	"rest_api_gin/internal/domains"

	"golang.org/x/crypto/bcrypt"
)

var _ domains.AuthService = (*AuthService)(nil)

type AuthService struct {
	userRepo   domains.UserRepo
	roleRepo   domains.RolesRepo
	jwtService domains.JWTService
}

func NewAuthService(userRepo domains.UserRepo,
	roleRepo domains.RolesRepo,
	jwtService domains.JWTService) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		roleRepo:   roleRepo,
		jwtService: jwtService,
	}
}

func (a *AuthService) LoginService(email string, password string) (string, error) {
	user, err := a.userRepo.GetByEmail(email)
	if err != nil {
		return "", fmt.Errorf("user not found %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("password incorrect %v", err)
	}

	roles, err := a.roleRepo.GetRoleByEmail(email)
	if err != nil {
		return "", fmt.Errorf("geting role failed %v", err)
	}

	token, err := a.jwtService.GenerateJWT(user.Email, roles)
	if err != nil {
		return "", fmt.Errorf("generate token failed %v", err)
	}

	return token, nil

}
