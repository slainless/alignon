package consumer

import "github.com/gofiber/fiber/v2"

func (s *Service) loan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("register")
	}
}
