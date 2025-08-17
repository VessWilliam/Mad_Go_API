package service

import (
	"fmt"
	"rest_api_gin/internal/claims"
	"rest_api_gin/internal/domains"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var _ domains.JWTService = (*JWTService)(nil)

type JWTService struct {
	secret   string
	duration time.Duration
}

func NewJWTService(secret string, duration time.Duration) *JWTService {
	if secret == "" || duration == 0 {
		panic("JWT secret empty or duration is 0")
	}

	return &JWTService{
		secret:   secret,
		duration: duration,
	}
}

func (j *JWTService) GenerateJWT(userId string, roles []string) (string, error) {

	claims := &claims.Claims{
		UserId: userId,
		Role:   roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTService) ValidateJWT(tokenString string) (*claims.Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &claims.Claims{}, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpectation signing method : %v", t.Header["alg"])
		}

		return []byte(j.secret), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token %v", err)
	}

	if claims, ok := parsedToken.Claims.(*claims.Claims); ok {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims

}
