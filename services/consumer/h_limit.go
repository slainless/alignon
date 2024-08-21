package consumer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slainless/my-alignon/pkg/platform"
)

func (s *Service) limit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		consumer := s.MustGetConsumer(c)
		limit, err := s.consumerManager.GetLimit(c.Context(), consumer.ID)
		if err != nil {
			if err != platform.ErrNoLimitSetYet {
				return c.Status(404).SendString(err.Error())
			}
			return c.Status(500).SendString("Fail to get limit")
		}

		return c.JSON(limit)
	}
}
