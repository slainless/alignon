package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/slainless/my-alignon/pkg/platform"
)

type EmailJWTAuthService struct {
	secret []byte
}

func NewEmailJWTAuthService(secret []byte) *EmailJWTAuthService {
	return &EmailJWTAuthService{
		secret: secret,
	}
}

// Credential implements platform.AuthService.
func (*EmailJWTAuthService) Credential(ctx *fiber.Ctx) (any, error) {
	token := ctx.Get("Authorization")
	if token == "" {
		return nil, platform.ErrEmptyCredential
	}

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
		return token, nil
	} else {
		return nil, platform.ErrUnsupportedCredential
	}
}

// ServiceID implements platform.AuthService.
func (*EmailJWTAuthService) ServiceID() string {
	return "supabase_jwt"
}

type Claims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

// Validate implements platform.AuthService.
func (s *EmailJWTAuthService) Validate(ctx context.Context, credential any) (email string, err error) {
	token, ok := credential.(string)
	if !ok {
		return "", platform.ErrInvalidCredential
	}

	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, platform.ErrUnsupportedHeader
		}
		return s.secret, nil
	})
	if err != nil {
		return "", errors.Join(err, platform.ErrInvalidCredential)
	}

	claim, ok := t.Claims.(*Claims)
	if !ok {
		return "", platform.ErrInvalidCredential
	}

	return claim.Email, nil
}
