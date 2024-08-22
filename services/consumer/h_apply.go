package consumer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slainless/my-alignon/internal/util"
	"github.com/slainless/my-alignon/internal/valid"
)

type ApplyPayload struct {
	Tenor    int16    `json:"tenor" validate:"required,min=1,max=4"`
	Catalogs []string `json:"catalogs" validate:"required,min=1,dive,required,uuid"`
}

func (s *Service) apply() fiber.Handler {
	return func(c *fiber.Ctx) error {
		consumer := s.MustGetConsumer(c)

		var payload ApplyPayload
		err := c.BodyParser(&payload)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}

		if err := valid.Struct(&payload); err != nil {
			return err
		}

		loan, err := s.loanManager.Apply(c.Context(), consumer, payload.Tenor, util.MustParseUUIDs(payload.Catalogs))
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(201).JSON(loan)
	}
}
