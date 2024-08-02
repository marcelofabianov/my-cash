package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/my-cash/pkg/logger"
)

func LoggingMiddleware(logger *logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		statusCode := c.Response().StatusCode()

		if statusCode >= 500 {
			logger.Error("request_error",
				logger.Field("method", c.Method()),
				logger.Field("path", c.Path()),
				logger.Field("remote_addr", c.IP()),
				logger.Field("duration", time.Since(start).String()),
				logger.Field("version", "v1"),
				logger.FieldInt("status", statusCode),
				logger.FieldError(err),
			)
			return err
		}

		if statusCode >= 400 && statusCode < 500 {
			logger.Error("request_error",
				logger.Field("method", c.Method()),
				logger.Field("path", c.Path()),
				logger.Field("remote_addr", c.IP()),
				logger.Field("duration", time.Since(start).String()),
				logger.Field("version", "v1"),
				logger.FieldInt("status", statusCode),
				logger.FieldError(err),
			)
			return err
		}

		logger.Info("request_completed",
			logger.Field("method", c.Method()),
			logger.Field("path", c.Path()),
			logger.Field("remote_addr", c.IP()),
			logger.Field("duration", time.Since(start).String()),
			logger.Field("version", "v1"),
			logger.FieldInt("status", c.Response().StatusCode()),
		)
		return err
	}
}
