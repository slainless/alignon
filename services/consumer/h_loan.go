package consumer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slainless/my-alignon/pkg/platform"
)

func (s *Service) loan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		consumer := s.MustGetConsumer(c)
		loan, err := s.loanManager.GetCurrentLoanOfConsumer(c.Context(), consumer.ID)
		if err != nil {
			if err == platform.ErrConsumerNotBorrowingAnyNow {
				return c.Status(404).SendString(err.Error())
			}
			return c.Status(500).SendString("Fail to get loan")
		}

		return c.JSON(loan)
	}
}
