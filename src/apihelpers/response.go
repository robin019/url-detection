package apihelpers

import "github.com/gofiber/fiber/v2"

func Success(ctx *fiber.Ctx, data interface{}) error {
	ctx.Status(200)
	return ctx.JSON(&fiber.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

func Failed(ctx *fiber.Ctx, apiError *ApiError) error {
	ctx.Status(apiError.Code)
	return ctx.JSON(&fiber.Map{
		"code":    apiError.Code,
		"message": apiError.Error.Error(),
		"data":    nil,
	})
}
