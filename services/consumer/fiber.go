package consumer

import "github.com/gofiber/fiber/v2"

func (s *Service) Mount(app fiber.Router) {
	// register to system, must be authenticated by 3rd party auth
	app.Post("/register", s.register())

	// patch user data
	app.Patch("/profile", s.profile())

	app.Get("/limit", s.limit())

	// get current loan detail
	app.Get("/loan", s.loan())
	// get current transaction detail
	app.Get("/loan/transaction/:id", s.transaction())

	// apply for loan
	app.Post("/apply", s.apply())

	// get loan history
	app.Get("/history", s.history())
}
