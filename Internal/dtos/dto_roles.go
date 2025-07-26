package dtos

// swagger:model RegisterRoleRequest
type RegisterRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

// swagger:model RegisterRoleResponse
type GetSingleRoleResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
