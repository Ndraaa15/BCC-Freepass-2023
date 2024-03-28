package logcustom

import (
	custom_error "bcc-freepass-2023/pkg/error"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ILogger interface {
	RequestLogger() fiber.Handler
	ErrorLogger(err custom_error.CustomError)
}

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() ILogger {
	newLogger := logrus.New()
	newLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	newLogger.SetOutput(os.Stdout)

	return &Logger{
		logger: newLogger,
	}
}

func (l *Logger) RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		l.logger.WithFields(logrus.Fields{
			"ip":     c.IP(),
			"method": c.Method(),
			"uri":    c.OriginalURL(),
		}).Info("Request")
		return c.Next()
	}
}

func (l *Logger) ErrorLogger(err custom_error.CustomError) {
	l.logger.WithFields(logrus.Fields{
		"code":     err.Code,
		"location": err.Location,
		"error":    err.Err,
	}).Error("Process")
}
