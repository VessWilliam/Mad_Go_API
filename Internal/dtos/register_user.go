package dtos

type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type GetAllUserResponse struct {
	UserList []GetSingleUserResponse `json:"Userlist"`
}

type GetSingleUserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
