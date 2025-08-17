package middleware

import (
	"context"
	"net/http"
	"rest_api_gin/internal/claims"
	"rest_api_gin/internal/dtos"
	"rest_api_gin/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService service.JWTService
}

func NewAuthMiddleware(jwtService service.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(string(claims.AuthorizationString))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtos.ErrorResponse{Message: "missing token"})
			return
		}

		parts := strings.Fields(authHeader)
		var tokenString string
		switch {
		case len(parts) == 2 && strings.EqualFold(parts[0], string(claims.BearerString)):
			tokenString = parts[1]
		case len(parts) == 1:
			tokenString = parts[0]
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtos.ErrorResponse{Message: "invalid auth header format"})
			return
		}

		parsedClaims, err := m.jwtService.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dtos.ErrorResponse{Message: "invalid token: " + err.Error()})
			return
		}

		ctx := context.WithValue(c.Request.Context(), claims.ClaimsContextKey, parsedClaims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
