package helper

import (
	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		ActualTag   string
		Params      string
		Value       interface{}
	}
)

type Validator struct {
	Validate *validator.Validate
}

//
//func NewValidator(validate *validator.Validate) Validator {
//	return Validator{Validate: validate}
//}
//
//func (v *Validator) validateStruct(payload interface{}) []ErrorResponse {
//	var validationErrors []ErrorResponse
//	errs := v.Validate.Struct(payload)
//	if errs != nil {
//		for _, err := range errs.(validator.ValidationErrors) {
//			var elem ErrorResponse
//
//			elem.FailedField = err.Field()   // Export struct field name
//			elem.ActualTag = err.ActualTag() // Export struct tag
//			elem.Value = err.Value()         // Export field value
//			elem.Params = err.Param()
//			elem.Error = true
//
//			validationErrors = append(validationErrors, elem)
//		}
//	}
//	return validationErrors
//}
//
//func (v *Validator) formFieldErrorMessage(fe ErrorResponse) string {
//	var sb strings.Builder
//
//	sb.WriteString("validation failed on field '" + fe.FailedField + "'")
//	sb.WriteString(", condition: " + fe.Tag)
//
//	if fe.Params() != "" {
//		sb.WriteString(" { " + fe.Params + " }")
//	}
//
//	if fe.Value() != nil && fe.Value() != "" {
//		sb.WriteString(fmt.Sprintf(", actual: %v", fe.Value()))
//	}
//
//	return sb.String()
//}
