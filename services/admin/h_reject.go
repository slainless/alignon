package admin

import "github.com/gofiber/fiber/v2"

func (s *Service) reject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("register")
	}
}
