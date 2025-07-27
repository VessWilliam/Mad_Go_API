package handler

import (
	"net/http"
	"rest_api_gin/internal/domains"
	"rest_api_gin/internal/dtos"
	"rest_api_gin/internal/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleHandle struct {
	RoleService *service.RoleService
}

func NewRoleHandle(s *service.RoleService) *RoleHandle {
	return &RoleHandle{RoleService: s}
}

// RegisterUser godoc
// @Summary      Register a new role
// @Description  Registers a role with name
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        role  body      dtos.RegisterRoleRequest  true  "Role registration data"
// @Success      201   {object}  dtos.GetSingleRoleResponse
// @Failure      400   {object}  dtos.ErrorResponse "Invalid input"
// @Failure      500   {object}  dtos.ErrorResponse "Could not register"
// @Router       /register_role [post]
func (h *RoleHandle) RegisterRole(c *gin.Context) {
	var req dtos.RegisterRoleRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	role := domains.Role{
		Name: strings.ToUpper(req.Name),
	}

	if err := h.RoleService.RegisterRoleService(&role); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dtos.GetSingleRoleResponse{
		Id:   role.Id,
		Name: role.Name,
	})
}

// GetAllUsers godoc
// @Summary      Get all roles
// @Description  Retrieve all registered roles
// @Tags         role
// @Accept       json
// @Produce      json
// @Success      200   {object}  dtos.GetAllRoleResponse
// @Failure      500   {object}  dtos.ErrorResponse "Internal server Error"
// @Router       /get_roles [get]
func (h *RoleHandle) GetRoles(c *gin.Context) {

	rolelist, err := h.RoleService.GetAllRoleService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	dtoList := make([]dtos.GetSingleRoleResponse, 0, len(rolelist))

	for _, role := range rolelist {
		dto := dtos.GetSingleRoleResponse{
			Id:   role.Id,
			Name: role.Name,
		}
		dtoList = append(dtoList, dto)
	}

	response := dtos.GetAllRoleResponse{
		RoleList: dtoList,
	}

	c.JSON(http.StatusOK, response)
}

// GetById godoc
// @Summary      Get role by ID
// @Description  Retrieve a single role by ID
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Role ID"
// @Success      200  {object}  dtos.GetSingleRoleResponse
// @Failure      500   {object}  dtos.ErrorResponse "Internal server Error"
// @Router       /getbyid_role/{id} [get]
func (h *RoleHandle) GetRolesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Invalid role ID"})
		return
	}

	role, err := h.RoleService.GetRoleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}
	response := dtos.GetSingleRoleResponse{
		Id:   role.Id,
		Name: role.Name,
	}

	c.JSON(http.StatusOK, response)
}

// GetById godoc
// @Summary      Delete role by ID
// @Description  Delete a single role by ID
// @Tags         role
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Role ID"
// @Success      200  {object}  dtos.GetSingleRoleResponse
// @Failure      500   {object}  dtos.ErrorResponse "Internal server Error"
// @Router       /deletebyid_role/{id} [delete]
func (h *RoleHandle) DeleteById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	role, err := h.RoleService.GetRoleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.RoleService.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, dtos.GetSingleRoleResponse{
		Id:   role.Id,
		Name: role.Name,
	})
}
