package platform

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slainless/my-alignon/internal/util"
)

type AuthManager struct {
	service      AuthService
	errorTracker ErrorTracker
}

func NewAuthManager(service AuthService, tracker ErrorTracker) *AuthManager {
	return &AuthManager{
		service:      service,
		errorTracker: tracker,
	}
}

func (m *AuthManager) Validate(c *fiber.Ctx) (email string, err error) {
	credential, err := m.service.Credential(c)
	if err != nil {
		switch {
		case util.IsCommonError(err, CommonAuthErrors):
		default:
			m.errorTracker.Report(c.Context(), err)
		}

		return "", err
	}

	email, err = m.service.Validate(c.Context(), credential)
	if err != nil {
		switch {
		case util.IsCommonError(err, CommonAuthErrors):
		default:
			m.errorTracker.Report(c.Context(), err)
		}
		return "", err
	}

	return email, nil
}
