package config

import (
	customErr "bcc-freepass-2023/pkg/error"
	"bcc-freepass-2023/pkg/response"
	"bcc-freepass-2023/pkg/validation"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func NewFiber() *fiber.App {
	return fiber.New(
		fiber.Config{
			EnablePrintRoutes: true,
			CaseSensitive:     true,
			StrictRouting:     true,
			ErrorHandler:      FiberErrorHandler(),
		},
	)
}

func FiberErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		if ce, ok := err.(*customErr.CustomError); ok {
			return response.Failed(c, ce.Code, ce.Message, ce.Err)
		}

		if iv, ok := err.(*validator.InvalidValidationError); ok {
			err := fmt.Sprintf("Invalid validation: %s", iv)
			return response.Failed(c, fiber.StatusBadRequest, "Invalid validation", err)
		}

		if fe, ok := err.(validator.ValidationErrors); ok {
			out := validation.GetValidationError(fe)
			return response.Failed(c, fiber.StatusBadRequest, "Invalid request", out)
		}

		return response.Failed(c, fiber.StatusInternalServerError, "Internal server error", err)
	}
}
