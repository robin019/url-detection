package middleware

import (
	"errors"

	"github.com/robin019/url-detection/src/utils/logger"

	"github.com/robin019/url-detection/src/apihelpers"

	"github.com/gofiber/fiber/v2"
)

// Fiber does not handle panics by default
// Use it as a middleware to catch panics
func Recover(ctx *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.ApiLog().Error(r)
			err = apihelpers.Failed(ctx, &apihelpers.ApiError{
				Code:  fiber.StatusInternalServerError,
				Error: errors.New("internal server error"),
			})
		}
	}()

	// Return err if exist, else move to next handler
	return ctx.Next()
}
