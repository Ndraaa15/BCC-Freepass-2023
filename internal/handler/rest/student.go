package rest

import (
	"bcc-freepass-2023/internal/usecase"
	"bcc-freepass-2023/model"
	customErr "bcc-freepass-2023/pkg/error"
	logcustom "bcc-freepass-2023/pkg/log"
	"bcc-freepass-2023/pkg/response"

	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	Identity       string
	studentUsecase usecase.IStudentUsecase
	logger         logcustom.ILogger
}

func NewStudentHandler(studentUsecase usecase.IStudentUsecase, logger logcustom.ILogger) *StudentHandler {
	return &StudentHandler{
		Identity:       "student",
		studentUsecase: studentUsecase,
	}
}

func (h *StudentHandler) MountEndpoint(fiber fiber.Router) {
	fiber.Post("/register", h.logger.RequestLogger(), h.StudentRegister)
}

func (h *StudentHandler) StudentRegister(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := *new(model.StudentRegister)

	if err := c.BodyParser(&req); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request", err)
	}

	resp, err := h.studentUsecase.CreateStudent(ctx, req)
	var customErr customErr.CustomError
	if errors.As(err, customErr) {
		h.logger.ErrorLogger(customErr)
		return response.Failed(c, customErr.Code, "Failed Create User", err)
	}

	return response.Success(c, fiber.StatusOK, "Success", resp)
}
