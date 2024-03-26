package rest

import (
	"bcc-freepass-2023/internal/usecase"
	"bcc-freepass-2023/model"
	"bcc-freepass-2023/pkg/log"
	"bcc-freepass-2023/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	Identity       string
	studentUsecase usecase.IStudentUsecase
}

func NewStudentHandler(studentUsecase usecase.IStudentUsecase) *StudentHandler {
	return &StudentHandler{
		Identity:       "student",
		studentUsecase: studentUsecase,
	}
}

func (h *StudentHandler) MountEndpoint(fiber fiber.Router, logger log.ILogger) {
	fiber.Post("/register", logger.RequestLogger(), h.StudentRegister)
}

func (h *StudentHandler) StudentRegister(c *fiber.Ctx) error {
	request := model.StudentRegister{}

	if err := c.BodyParser(&request); err != nil {
		return response.Failed(c, fiber.StatusBadRequest, "Invalid request", err)
	}

	return response.Success(c, fiber.StatusOK, "Success", nil)
}
