package handler

import (
	"net/http"
	"rest_api_gin/internal/domains"
	"rest_api_gin/internal/dtos"
	"rest_api_gin/internal/service"
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

	role := domains.Roles{
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
