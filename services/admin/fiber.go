package admin

import "github.com/gofiber/fiber/v2"

func (s *Service) Mount(app fiber.Router) {
	// get loans
	// by default skipping fully paid loan
	app.Get("/loan", s.loan())

	// get loan data
	app.Get("/loan/:id", s.loan_detail())
	// approve
	app.Post("/loan/:id/approve", s.approve())
	// reject
	app.Post("/loan/:id/reject", s.reject())
	// paid loan
	app.Post("/loan/:id/paid", s.paid())

	// get customers
	app.Get("/customer", s.customers())
	// get customer detail
	app.Get("/customer/:id", s.customer_detail())
}
