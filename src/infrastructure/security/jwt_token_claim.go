package security

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenClaim struct {
	jwt.RegisteredClaims
	Payload []byte `json:"payload"`
}
