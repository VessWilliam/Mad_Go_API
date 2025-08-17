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

// RegisterRole godoc
// @Summary      Register a new role
// @Description  Registers a role with name
// @Tags         roles
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

// GetRoles godoc
// @Summary      Get all roles
// @Description  Retrieve all registered roles
// @Tags         roles
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

// GetRolesById godoc
// @Summary      Get role by ID
// @Description  Retrieve a single role by ID
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Role ID"
// @Success      200  {object}  dtos.GetSingleRoleResponse
// @Failure      500   {object}  dtos.ErrorResponse "Internal server Error"
// @Security     BearerAuth
// @Router       /getbyid_role/{id} [get]
func (h *RoleHandle) GetRolesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "invalid role id"})
		return
	}

	role, err := h.RoleService.GetRoleByIdService(id)
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

// DeleteById godoc
// @Summary      Delete role by ID
// @Description  Delete a single role by ID
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Role ID"
// @Success      200  {object}  dtos.SuccessResponse
// @Failure      500   {object}  dtos.ErrorResponse "Internal server Error"
// @Security     BearerAuth
// @Router       /deletebyid_role/{id} [delete]
func (h *RoleHandle) DeleteById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.RoleService.GetRoleByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	err = h.RoleService.DeleteByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, dtos.SuccessResponse{Message: "deleted role success !"})
}

// UpdateRole godoc
// @Summary      Update Role
// @Description  Update role body
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        request body      dtos.UpdateRoleRequest  true  "Role Body"
// @Success      200   {object}  dtos.SuccessResponse
// @Failure      400   {object}  dtos.ErrorResponse
// @Failure      500   {object}  dtos.ErrorResponse
// @Security     BearerAuth
// @Router       /update_role [put]
func (h *RoleHandle) UpdateRole(c *gin.Context) {
	var req dtos.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Invalid JSON format"})
		return
	}

	role := domains.Role{
		Id:   req.Id,
		Name: req.Name,
	}

	if err := h.RoleService.UpdateRoleService(&role); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResponse{Message: "role successfully updated"})
}
