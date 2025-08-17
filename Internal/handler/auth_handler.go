package handler

import (
	"net/http"
	"rest_api_gin/internal/dtos"
	"rest_api_gin/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandle(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login godoc
// @Summary      Login
// @Description  Authenticate user and return JWT Token
// @Tags         auth login
// @Accept       json
// @Produce      json
// @Param        login  body      dtos.LoginRequest  true  "Login Request"
// @Success      200    {object}  dtos.TokenResponse
// @Failure      400    {object}  dtos.ErrorResponse "Invalid input"
// @Failure      401    {object}  dtos.ErrorResponse "Unauthorized"
// @Failure      500    {object}  dtos.ErrorResponse "Internal server Error"
// @Router       /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dtos.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	JWTtoken, err := h.authService.LoginService(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dtos.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.TokenResponse{JWTToken: JWTtoken})
}
