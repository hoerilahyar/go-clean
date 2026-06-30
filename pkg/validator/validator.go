package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate validates a struct and returns a human-readable error.
func Validate(v any) error {

	if err := validate.Struct(v); err != nil {

		var validationErrors validator.ValidationErrors

		if errors.As(err, &validationErrors) {

			messages := make([]string, 0, len(validationErrors))

			for _, field := range validationErrors {

				messages = append(
					messages,
					buildMessage(field),
				)
			}

			return errors.New(strings.Join(messages, ", "))
		}

		return err
	}

	return nil
}

func buildMessage(
	field validator.FieldError,
) string {

	name := field.Field()

	switch field.Tag() {

	case "required":
		return fmt.Sprintf("%s is required", name)

	case "email":
		return fmt.Sprintf("%s must be a valid email address", name)

	case "min":
		return fmt.Sprintf("%s must be at least %s characters", name, field.Param())

	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", name, field.Param())

	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", name, field.Param())

	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", name, field.Param())

	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", name, field.Param())

	default:
		return fmt.Sprintf("%s is invalid", name)
	}
}
