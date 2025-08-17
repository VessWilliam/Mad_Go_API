package domains

import "rest_api_gin/internal/claims"

type JWTService interface {
	GenerateJWT(userId string, roles []string) (string, error)
	ValidateJWT(tokenString string) (*claims.Claims, error)
}
