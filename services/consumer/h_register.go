package consumer

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/slainless/my-alignon/internal/util"
	"github.com/slainless/my-alignon/internal/valid"
	"github.com/slainless/my-alignon/pkg/platform"
	"github.com/valyala/fasthttp"
)

func (s *Service) register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email, err := s.authManager.Validate(c)
		if err != nil {
			if util.IsCommonError(err, platform.CommonAuthErrors) {
				return c.Status(400).SendString(err.Error())
			}
			return c.Status(500).SendString("Fail to validate token")
		}

		var payload platform.ConsumerRegisterInput
		err = c.BodyParser(&payload)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}

		if err := valid.Struct(&payload); err != nil {
			return err
		}

		payload.BirthDate, err = time.Parse(time.DateOnly, payload.P_BirthDate)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		payload.Email = email

		ktp, err := c.FormFile("ktp_photo")
		if err != nil {
			if err == fasthttp.ErrMissingFile {
				return c.Status(400).SendString("ktp_photo is required")
			}
			return c.Status(500).SendString(err.Error())
		}

		selfie, err := c.FormFile("selfie_photo")
		if err != nil {
			if err == fasthttp.ErrMissingFile {
				return c.Status(400).SendString("selfie_photo is required")
			}
			return c.Status(500).SendString(err.Error())
		}

		consumer, err := s.consumerManager.Register(c.Context(), &payload, ktp, selfie)
		if err != nil {
			return c.Status(500).SendString("Fail to register user")
		}
		return c.Status(201).JSON(consumer)
	}
}
