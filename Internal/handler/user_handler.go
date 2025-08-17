package handler

import (
	"net/http"
	"rest_api_gin/internal/domains"
	"rest_api_gin/internal/dtos"
	"rest_api_gin/internal/service"
	"strconv"

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
// @Router       /register_user [post]
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

	if err := h.UserService.RegisterUserService(&user); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dtos.GetSingleUserResponse{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	})

}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve all registered users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  dtos.GetAllUserResponse
// @Failure      500   {object}  dtos.ErrorResponse "Internal server Error"
// @Router       /get_users [get]
func (h *UserHandle) GetUsers(c *gin.Context) {

	userlist, err := h.UserService.GetAllUserService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	dtoList := make([]dtos.GetSingleUserResponse, 0, len(userlist))
	for _, user := range userlist {

		rolelist := make([]dtos.RoleList, 0, len(userlist))
		for _, role := range user.Role {
			roleDto := dtos.RoleList{
				Name: role.Name,
			}
			rolelist = append(rolelist, roleDto)
		}

		dto := dtos.GetSingleUserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Roles: rolelist,
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
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  dtos.GetSingleUserResponse
// @Failure      500  {object}  dtos.ErrorResponse "Internal server Error"
// @Security     BearerAuth
// @Router       /getbyid_user/{id} [get]
func (h *UserHandle) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Invalid user ID"})
		return
	}

	user, err := h.UserService.GetUserByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	roleList := make([]dtos.RoleList, 0, len(user.Role))
	for _, role := range user.Role {
		dto := dtos.RoleList{
			Name: role.Name,
		}
		roleList = append(roleList, dto)
	}

	response := dtos.GetSingleUserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Roles: roleList,
	}

	c.JSON(http.StatusOK, response)
}

// Assign Role To User godoc
// @Summary      Assign Roles to User
// @Description  Assign multiple roles to a specific user by body
// @Tags         assign role
// @Accept       json
// @Produce      json
// @Param        request body      dtos.AssignRolesRequest  true  "User ID and Role IDs"
// @Success      200   {object}  dtos.SuccessResponse
// @Failure      400   {object}  dtos.ErrorResponse
// @Failure      500   {object}  dtos.ErrorResponse
// @Security     BearerAuth
// @Router       /assign-roles [put]
func (h *UserHandle) AssignRolesToUser(c *gin.Context) {
	var req dtos.AssignRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Invalid JSON format"})
		return
	}

	err := h.UserService.AssignRolesToUserService(req.UserId, req.RoleIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResponse{Message: "Roles successfully assigned"})
}

// Update User godoc
// @Summary      Update User
// @Description  Update user body
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body      dtos.UpdateUserRequest  true  "User Body"
// @Success      200   {object}  dtos.SuccessResponse
// @Failure      400   {object}  dtos.ErrorResponse
// @Failure      500   {object}  dtos.ErrorResponse
// @Security     BearerAuth
// @Router       /update_user [put]
func (h *UserHandle) UpdateUser(c *gin.Context) {
	var req dtos.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Invalid JSON format"})
		return
	}

	user := domains.User{
		Id:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.UserService.UpdateUserService(&user); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResponse{Message: "User successfully Updated"})
}
