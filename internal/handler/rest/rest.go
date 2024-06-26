package rest

import (
	"bcc-freepass-2023/internal/repository"
	"bcc-freepass-2023/internal/usecase"
	"bcc-freepass-2023/pkg/config"
	"bcc-freepass-2023/pkg/database/postgresql"
	errcustom "bcc-freepass-2023/pkg/error"
	logcustom "bcc-freepass-2023/pkg/log"
	"bcc-freepass-2023/pkg/validation"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	Success = iota
	BadConfig
	FailedRunServer
)

type Handler interface {
	MountEndpoint(fiber fiber.Router)
}

type Rest struct {
	fiber    *fiber.App
	handlers map[string]Handler
	logger   logcustom.ILogger
}

func InitializeServer() *Rest {
	var rest Rest
	rest.fiber = config.NewFiber()
	rest.handlers = make(map[string]Handler)
	rest.logger = logcustom.NewLogger()

	return &rest
}

func (r *Rest) registerHandler(name string, handler Handler) {
	r.handlers[name] = handler
}

func (r *Rest) newServer() error {
	if err := config.LoadConfig(); err != nil {
		return errcustom.NewCustomError(http.StatusInternalServerError, "[newServer] : load config", "failed to load config", err)
	}

	dbConn, err := postgresql.InitPostgreSQL()
	if err != nil {
		return errcustom.NewCustomError(http.StatusInternalServerError, "[newServer] : init postgresql", "failed to initialize postgresql", err)
	}

	validator := validation.NewValidator()

	repository := repository.New(dbConn)
	studentUsecase := usecase.NewStudentUsecase(repository)
	studentHandler := NewStudentHandler(studentUsecase, r.logger, validator)

	r.registerHandler(studentHandler.Identity, studentHandler)

	return nil
}

func (r *Rest) RunServer() (int, error) {
	if err := r.newServer(); err != nil {
		return BadConfig, err
	}

	for key := range r.handlers {
		routerGroup := r.fiber.Group(fmt.Sprintf("%s/%s/%s", "api", "v1", key))
		r.handlers[key].MountEndpoint(routerGroup)
	}

	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	if err := r.fiber.Listen(fmt.Sprintf("%s:%s", addr, port)); err != nil {
		return FailedRunServer, errcustom.NewCustomError(http.StatusInternalServerError, "[RunServer] : listen address and port", "failed to run server", err)
	}

	return Success, nil
}
