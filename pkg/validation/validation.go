package validation

import (
	"bcc-freepass-2023/model"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type validationError struct {
	Namespace       string `json:"namespace"`
	Field           string `json:"field"`
	StructNamespace string `json:"structNamespace"`
	Tag             string `json:"tag"`
	Message         string `json:"message"`
}

type IValidator interface {
	ValidateStruct(request any) error
}

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	v := &Validator{
		validator: validator.New(),
	}
	v.addValidation()
	return v
}

func (v *Validator) ValidateStruct(request any) error {
	if err := v.validator.Struct(request); err != nil {
		return err
	}
	return nil

}

func (v *Validator) addValidation() {
	v.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	v.validator.RegisterStructValidation(registerRequestValidation, model.StudentRegister{})
}

func GetValidationError(fe validator.ValidationErrors) any {
	out := make(map[string]interface{})

	for _, err := range fe {
		e := validationError{
			Namespace:       err.Namespace(),
			Field:           err.Field(),
			StructNamespace: err.StructNamespace(),
			Tag:             err.Tag(),
			Message:         getErrorMessage(err),
		}

		out[err.Field()] = e
	}

	return out
}

func registerRequestValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(model.StudentRegister)

	if len(user.Email) == 0 {
		sl.ReportError(user.Email, "email", "Email", "required", "")
	}
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s is not a valid email", fe.Field())
	case "max":
		return fmt.Sprintf("%s must be less than %s", fe.Field(), fe.Param())
	case "min":
		return fmt.Sprintf("%s must be more than %s", fe.Field(), fe.Param())
	case "number":
		return fmt.Sprintf("%s must be a number", fe.Field())
	case "e164":
		return fmt.Sprintf("%s must be a valid phone number", fe.Field())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}
