package middleware

import (
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/robin019/url-detection/src/utils/logger"
)

func RequestLog() fiber.Handler {
	return fiberlog.New(fiberlog.Config{
		TimeFormat: "2006-01-02 15:04:05",
		Format:     "[${time}] ${header:X-Real-IP} - ${status} - ${latency} ${method} ${path}\n",
		Output:     logger.RequestLog,
	})
}
