package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Status struct {
	IsSuccess bool `json:"isSuccess"`
	Code      int  `json:"code"`
}

func Success(c *fiber.Ctx, code int, message string, data any) error {
	return c.JSON(&Response{
		Status: Status{
			IsSuccess: true,
			Code:      code,
		},
		Message: message,
		Data:    data,
	}, fiber.MIMEApplicationJSON)
}

func Failed(c *fiber.Ctx, code int, message string, err any) error {
	return c.JSON(&Response{
		Status: Status{
			IsSuccess: false,
			Code:      code,
		},
		Message: message,
		Data:    err,
	}, fiber.MIMEApplicationJSON)
}
