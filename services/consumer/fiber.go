package consumer

import "github.com/gofiber/fiber/v2"

func (s *Service) Mount(app fiber.Router) {
	// register to system, must be authenticated by 3rd party auth
	app.Post("/register", s.register())

	protected := app.Group("/")
	protected.Use(s.auth())

	// patch user data
	protected.Patch("/profile", s.profile())

	protected.Get("/limit", s.limit())

	// get current loan detail
	protected.Get("/loan", s.loan())
	// get current transaction detail
	protected.Get("/loan/transaction/:id", s.transaction())

	// apply for loan
	protected.Post("/apply", s.apply())

	// get loan history
	protected.Get("/history", s.history())
}
