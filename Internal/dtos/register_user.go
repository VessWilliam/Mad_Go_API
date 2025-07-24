package dtos

import "rest_api_gin/internal/domains"

type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type GetAllUserResponse struct {
	UserList []*domains.User `json:"Userlist"`
}
