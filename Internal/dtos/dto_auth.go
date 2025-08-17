package dtos

// swagger:model LoginRequest
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// swagger:model LoginRequest
type TokenResponse struct {
	JWTToken string `json:"JWTtoken"`
}
