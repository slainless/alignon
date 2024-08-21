package consumer

// func (s *Service) transaction() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		consumer := s.MustGetConsumer(c)

// 		loan, err := s.loanManager.GetTransactionOfCurrentLoan(c.Context(), consumer.ID, transactionID)
// 		if err != nil {
// 			if err == platform.ErrConsumerNotBorrowingAnyNow {
// 				return c.Status(404).SendString(err.Error())
// 			}
// 			return c.Status(500).SendString("Fail to get transactions")
// 		}

// 		return c.JSON(loan)
// 	}
// }
