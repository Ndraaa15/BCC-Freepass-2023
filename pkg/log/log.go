package log

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type InterfaceLogger interface {
	RequestLogger() fiber.Handler
}

type Logger struct {
	logger *logrus.Logger
}

func NewLog() InterfaceLogger {
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
			"ip":        c.IP(),
			"method":    c.Method(),
			"uri":       c.OriginalURL(),
			"userAgent": c.Context().UserAgent(),
		}).Info("Request")
		return c.Next()
	}
}
