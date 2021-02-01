package service

import (
	"reflect"

	"github.com/robin019/url-detection/src/apihelpers"
)

func UrlCheck(params interface{}) (map[string]interface{}, *apihelpers.ApiError) {
	value := reflect.ValueOf(params).Elem()
	url := value.FieldByName("Url").String()
	return map[string]interface{}{
		"url":   url,
		"hello": "world",
	}, nil
}
