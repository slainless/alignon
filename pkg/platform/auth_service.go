package platform

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrUnsupportedCredential = errors.New("unsupported credential")
	ErrEmptyCredential       = errors.New("empty credential")
	ErrInvalidCredential     = errors.New("invalid credential")
	ErrUnsupportedHeader     = errors.New("unsupported header")
)

var CommonAuthErrors = []error{
	ErrUnsupportedCredential,
	ErrEmptyCredential,
	ErrInvalidCredential,
	ErrUnsupportedHeader,
}

type AuthService interface {
	Credential(c *fiber.Ctx) (credential any, err error)
	Validate(ctx context.Context, credential any) (email string, err error)
}
