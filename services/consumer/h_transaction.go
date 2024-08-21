package consumer

import "github.com/gofiber/fiber/v2"

func (s *Service) transaction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("register")
	}
}
