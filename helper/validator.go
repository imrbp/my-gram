package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Validator struct {
	Validate *validator.Validate
}

func NewValidator(validate *validator.Validate) Validator {
	return Validator{Validate: validate}
}

func (v *Validator) formFieldErrorMessage(fe validator.FieldError) string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + fe.Field() + "'")
	sb.WriteString(", condition: " + fe.ActualTag())

	if fe.Param() != "" {
		sb.WriteString(" { " + fe.Param() + " }")
	}

	if fe.Value() != nil && fe.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", fe.Value()))
	}

	return sb.String()
}

func (v *Validator) validateStruct(payload interface{}) error {
	err := v.Validate.Struct(payload)
	errMessage := ""
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errMessage += v.formFieldErrorMessage(err) + "\n"
		}
		return fiber.NewError(fiber.StatusBadRequest, errMessage)
	} else {
		return nil
	}
}

func (v *Validator) ParseBody(ctx *fiber.Ctx, payload interface{}) error {
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}
	return v.validateStruct(payload)

}
