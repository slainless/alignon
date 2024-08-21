package consumer

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Service) history() fiber.Handler {
	return func(c *fiber.Ctx) error {
		consumer := s.MustGetConsumer(c)
		loans, err := s.loanManager.GetLoans(c.Context(), consumer.ID)
		if err != nil {
			return c.Status(500).SendString("Fail to get loans")
		}

		return c.JSON(loans)
	}
}
