package consumer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slainless/my-alignon/internal/util"
	"github.com/slainless/my-alignon/pkg/platform"
)

func (s *Service) auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email, err := s.authManager.Validate(c)
		if err != nil {
			if util.IsCommonError(err, platform.CommonAuthErrors) {
				return c.Status(401).SendString(err.Error())
			}
			return c.Status(500).SendString("Fail to validate user")
		}

		consumer, err := s.consumerManager.GetByEmail(c.Context(), email)
		if err != nil {
			if err == platform.ErrConsumerNotFound {
				return c.Status(401).SendString(err.Error())
			}
			return c.Status(500).SendString("Fail to get user")
		}

		c.Locals(LocalsKeyConsumer, consumer)
		return c.Next()
	}
}

type LocalsKey string

const LocalsKeyConsumer = LocalsKey("consumer")

func (s *Service) MustGetConsumer(c *fiber.Ctx) *platform.Consumer {
	if consumer, ok := c.Locals(LocalsKeyConsumer).(*platform.Consumer); ok {
		return consumer
	} else {
		panic("consumer not found")
	}
}
