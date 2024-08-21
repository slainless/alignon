package valid

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

var ErrInvalidInput = fiber.NewError(500, "Invalid input received in validator, high chance of implementation error")

func Struct(data any) error {
	err := Validator.Struct(data)
	if err == nil {
		return nil
	}

	if err := err.(*validator.InvalidValidationError); err != nil {
		return ErrInvalidInput
	}

	if err := err.(validator.ValidationErrors); err != nil {
		msgs := make([]string, 0, len(err))
		for _, err := range err {
			msgs = append(msgs, fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", err.Field(), err.Value(), err.Tag()))
		}

		return fiber.NewError(400, "Invalid input:\n"+strings.Join(msgs, "\n"))
	}

	return fiber.NewError(500, err.Error())
}
