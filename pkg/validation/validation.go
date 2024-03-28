package validation

import (
	"bcc-freepass-2023/model"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type validationError struct {
	Namespace       string `json:"namespace"`
	Field           string `json:"field"`
	StructNamespace string `json:"structNamespace"`
	StructField     string `json:"structField"`
	Tag             string `json:"tag"`
	ActualTag       string `json:"actualTag"`
	Message         string `json:"message"`
}

type IValidator interface {
	ValidateStruct(request any) any
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

func (v *Validator) ValidateStruct(request any) any {
	if err := v.validator.Struct(request); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}

		out := make(map[string]interface{})

		for _, err := range err.(validator.ValidationErrors) {
			e := validationError{
				Namespace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Message:         getErrorMessage(err),
			}

			out[err.Field()] = e

			indent, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			fmt.Println(string(indent))
		}

		return out
	}
	return nil
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
