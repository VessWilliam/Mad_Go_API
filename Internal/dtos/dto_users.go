package dtos

// swagger:model RegisterUserRequest
type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

// swagger:model GetAllUserResponse
type GetAllUserResponse struct {
	UserList []GetSingleUserResponse `json:"Userlist"`
}

// swagger:model GetSingleUserResponse
type GetSingleUserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
