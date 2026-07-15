package security

import (
	"errors"
	"example.com/PROJECT_NAME/src/core/application/response"
	"example.com/PROJECT_NAME/src/core/domain/ports/security"
	"example.com/PROJECT_NAME/src/core/utils"
	"example.com/PROJECT_NAME/src/infrastructure/env"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtTokenAdapter struct {
	secret []byte
}

// @inject
func NewJwtTokenAdapter(cfg *env.Env) security.JwtTokenProvider {
	return &JwtTokenAdapter{secret: []byte(cfg.JwtSecret)}
}

func (this *JwtTokenAdapter) Encode(data *security.JwtToken) string {
	claim := &JwtTokenClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   data.Subject,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(data.Exp),
			ID:        uuid.New().String(),
		},
		Payload: data.Payload,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := t.SignedString(this.secret)
	if err != nil {
		return ""
	}
	return token
}

func (this *JwtTokenAdapter) Decode(tokenString string) (*security.JwtToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtTokenClaim{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, response.NewSafeError(response.CodeInvalidToken, errors.New("invalid token signing method"), utils.CallerPath(1))
		}
		return this.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, response.NewSafeError(response.CodeExpiredToken, err, utils.CallerPath(1))
		}
		return nil, response.NewSafeError(response.CodeInvalidToken, err, utils.CallerPath(1))
	}

	claims, ok := token.Claims.(*JwtTokenClaim)
	if !ok || !token.Valid {
		return nil, response.NewSafeError(response.CodeInvalidToken, err, utils.CallerPath(1))
	}

	return &security.JwtToken{
		Exp:     claims.ExpiresAt.Time,
		Subject: claims.Subject,
		Payload: claims.Payload,
	}, nil
}
