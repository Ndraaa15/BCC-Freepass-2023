package rest

import "github.com/gofiber/fiber/v2"

type Rest struct {
	fiber fiber.Router
}

func NewRest() *Rest {
	return &Rest{
		fiber: fiber.New(
			fiber.Config{
				EnablePrintRoutes: true,
				CaseSensitive:     true,
				StrictRouting:     true,
			},
		),
	}
}

func (h *Rest) GetRouter() {
	h.fiber.Get("/", h.StudentLogin)
}
