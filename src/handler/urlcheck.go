package handler

import (
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/robin019/url-detection/src/apihelpers"
	"github.com/robin019/url-detection/src/service"
)

//full compare the passed in URL with the URL in database and return all websites that consider the url malicious.
func UrlCheck(ctx *fiber.Ctx) error {
	// Unescape a query string to avoid any ambiguity
	// e.g., /malicious_url?url=https://some-url.com and /malicious_url?url=https://some-url.com?detail=1
	urlDecode, err := url.QueryUnescape(ctx.Query("url"))

	// url.QueryUnescape() returns an error if any % is not followed by two hexadecimal digits.
	if err != nil {
		return apihelpers.Failed(ctx, &apihelpers.ApiError{
			Code:  fiber.StatusBadRequest,
			Error: err,
		})
	}

	type rule struct {
		Url string `validate:"required,url"`
	}
	params := &rule{
		Url: urlDecode,
	}

	err = validator.New().Struct(params)
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
