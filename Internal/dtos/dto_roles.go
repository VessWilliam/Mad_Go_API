package dtos

// swagger:model RegisterRoleRequest
type RegisterRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

// swagger:model GetAllRoleResponse
type GetAllRoleResponse struct {
	RoleList []GetSingleRoleResponse `json:"Rolelist"`
}

// swagger:model RegisterRoleResponse
type GetSingleRoleResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// swagger:model UpdateRoleRequest
type UpdateRoleRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
