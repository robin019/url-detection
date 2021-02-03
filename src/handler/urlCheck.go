package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/robin019/url-detection/src/apihelpers"
	"github.com/robin019/url-detection/src/service"
)

//full compare the passed in URL with the URL in database and return all websites that consider the url malicious.
func UrlCheck(ctx *fiber.Ctx) error {
	type rule struct {
		Url string `validate:"required,url"`
	}
	params := &rule{
		Url: ctx.Query("url"),
	}

	err := validator.New().Struct(params)
	if err != nil {
		return apihelpers.Failed(ctx, &apihelpers.ApiError{
			Code:  fiber.StatusBadRequest,
			Error: err,
		})
	}

	if result, apiError := service.UrlCheck(params); apiError != nil {
		return apihelpers.Failed(ctx, apiError)
	} else {
		return apihelpers.Success(ctx, result)
	}
}
