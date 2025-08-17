package claims

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserId string   `json:"user_id"`
	Role   []string `json:"role"`
	jwt.RegisteredClaims
}

type ctxKey string
type authString string
type bearerString string

const (
	ClaimsContextKey    ctxKey       = "claims"
	AuthorizationString authString   = "Authorization"
	BearerString        bearerString = "Bearer"
)
