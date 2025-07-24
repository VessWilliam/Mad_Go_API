package handler

import (
	"net/http"
	"rest_api_gin/internal/domains"
	"rest_api_gin/internal/dtos"
	"rest_api_gin/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandle struct {
	UserService *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandle {
	return &UserHandle{UserService: s}
}

func (h *UserHandle) RegisterUser(c *gin.Context) {

	var req dtos.RegisterUserRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user := domains.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	if err := h.UserService.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.Id})

}

func (h *UserHandle) GetUser(c *gin.Context) {
	userlist, err := h.UserService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dtos.GetAllUserResponse{
		UserList: userlist,
	}

	c.JSON(http.StatusOK, response)

}
