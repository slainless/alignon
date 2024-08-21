package consumer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slainless/my-alignon/internal/util"
	"github.com/slainless/my-alignon/pkg/platform"
)

type RegisterPayload struct {
	Token string `json:"token" form:"token"`
}

func (s *Service) register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload RegisterPayload
		err := c.BodyParser(&payload)
		if err != nil {
			return err
		}

		email, err := s.authManager.Validate(c)
		if err != nil {
			if util.IsCommonError(err, platform.CommonAuthErrors) {
				return c.Status(400).SendString(err.Error())
			}
			return c.Status(500).SendString("Fail to validate token")
		}

		err = s.consumerManager.Register(c.Context(), email)
		if err != nil {
			return c.Status(500).SendString("Fail to register user")
		}
		return c.SendStatus(201)
	}
}
