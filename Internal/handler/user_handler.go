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

// RegisterUser godoc
// @Summary      Register a new user
// @Description  Registers a user with email, password, and name
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dtos.RegisterUserRequest  true  "User registration data"
// @Success      201   {object}  dtos.GetSingleUserResponse
// @Failure      400   {object}  dtos.ErrorResponse "Invalid input"
// @Failure      500   {object}  dtos.ErrorResponse "Could not register"
// @Router       /register [post]
func (h *UserHandle) RegisterUser(c *gin.Context) {

	var req dtos.RegisterUserRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	user := domains.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	if err := h.UserService.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.Id})

}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve all registered users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  dtos.GetAllUserResponse
// @Failure      500   {object}  dtos.ErrorResponse "Internal server Error"
// @Router       /getall [get]
func (h *UserHandle) GetUser(c *gin.Context) {

	userlist, err := h.UserService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	dtoList := make([]dtos.GetSingleUserResponse, 0, len(userlist))

	for _, user := range userlist {
		dto := dtos.GetSingleUserResponse{
			Email: user.Email,
			Name:  user.Name,
		}
		dtoList = append(dtoList, dto)
	}

	response := dtos.GetAllUserResponse{
		UserList: dtoList,
	}

	c.JSON(http.StatusOK, response)
}

// GetById godoc
// @Summary      Get user by ID
// @Description  Retrieve a single user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  dtos.GetSingleUserResponse
// @Failure      500   {object}  dtos.ErrorResponse "Internal server Error"
// @Router       /getbyid/{id} [get]
func (h *UserHandle) GetById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UserService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	response := dtos.GetSingleUserResponse{
		Email: user.Email,
		Name:  user.Name,
	}

	c.JSON(http.StatusOK, response)
}
