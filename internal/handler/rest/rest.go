package rest

import (
	"bcc-freepass-2023/internal/repository"
	"bcc-freepass-2023/internal/usecase"
	"bcc-freepass-2023/pkg/config"
	"bcc-freepass-2023/pkg/database/postgresql"
	"bcc-freepass-2023/pkg/log"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	Success = iota
	BadConfig
	FailedRunServer
)

type Handler interface {
	MountEndpoint(fiber fiber.Router, logger log.ILogger)
}

type Rest struct {
	fiber    *fiber.App
	handlers map[string]Handler
	logger   log.ILogger
}

func InitializeServer() *Rest {
	var rest Rest
	rest.fiber = fiber.New(
		fiber.Config{
			EnablePrintRoutes: true,
			CaseSensitive:     true,
			StrictRouting:     true,
		},
	)
	rest.handlers = make(map[string]Handler)
	rest.logger = log.NewLogger()
	return &rest
}

func (r *Rest) registerHandler(name string, handler Handler) {
	r.handlers[name] = handler
}

func (r *Rest) newServer() error {
	if err := config.LoadConfig(); err != nil {
		return err
	}

	dbConn, err := postgresql.InitPostgreSQL()
	if err != nil {
		return err
	}

	repository := repository.New(dbConn)
	studentUsecase := usecase.NewStudentUsecase(repository)
	studentHandler := NewStudentHandler(studentUsecase)
	r.registerHandler(studentHandler.Identity, studentHandler)

	return nil
}

func (r *Rest) RunServer() (int, error) {
	if err := r.newServer(); err != nil {
		return BadConfig, err
	}

	for key, handler := range r.handlers {
		routerGroup := r.fiber.Group(fmt.Sprintf("%s/%s/%s", "api", "v1", key))
		handler.MountEndpoint(routerGroup, r.logger)
	}

	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	if err := r.fiber.Listen(fmt.Sprintf("%s:%s", addr, port)); err != nil {
		return FailedRunServer, err
	}

	return Success, nil
}
