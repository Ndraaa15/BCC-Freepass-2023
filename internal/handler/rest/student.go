package rest

import (
	"bcc-freepass-2023/internal/usecase"
	"bcc-freepass-2023/model"
	logcustom "bcc-freepass-2023/pkg/log"
	"bcc-freepass-2023/pkg/response"
	"bcc-freepass-2023/pkg/validation"

	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	Identity       string
	studentUsecase usecase.IStudentUsecase
	logger         logcustom.ILogger
	validator      validation.IValidator
}

func NewStudentHandler(studentUsecase usecase.IStudentUsecase, logger logcustom.ILogger, validator validation.IValidator) *StudentHandler {
	return &StudentHandler{
		Identity:       "student",
		studentUsecase: studentUsecase,
		logger:         logger,
		validator:      validator,
	}
}

func (h *StudentHandler) MountEndpoint(fiber fiber.Router) {
	fiber.Post("/register", h.logger.RequestLogger(), h.StudentRegister)
}

func (h *StudentHandler) StudentRegister(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := model.StudentRegister{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validator.ValidateStruct(req); err != nil {
		return err
	}

	resp, err := h.studentUsecase.CreateStudent(ctx, req)
	if err != nil {
		return err
	}

	return response.Success(c, fiber.StatusCreated, "Success", resp)
}
