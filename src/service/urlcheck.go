package service

import (
	"errors"
	"reflect"
	"time"

	"github.com/robin019/url-detection/src/utils/logger"

	"github.com/gofiber/fiber/v2"

	"github.com/robin019/url-detection/src/psersistence/sql"

	"github.com/robin019/url-detection/src/apihelpers"
)

func UrlCheck(params interface{}) (result []map[string]interface{}, apiError *apihelpers.ApiError) {
	value := reflect.ValueOf(params).Elem()
	url := value.FieldByName("Url").String()

	// crop the last character if it is a '/'
	// e.g., http://google.com/ to http://google.com
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}
	type queryResult struct {
		Source           string    `gorm:"source"`
		SourceId         string    `gorm:"sourceId"`
		VerificationTime time.Time `gorm:"verification_time"`
	}
	var queryResults []*queryResult

	db := sql.DB()
	err := db.Table("malicious_url as mu").
		Select("md.source, md.source_id, md.verification_time").
		Joins("inner join malicious_url_detail AS md on mu.id = md.url_id").
		Where("md5(lower(mu.url)) = md5(lower(?))", url).
		Scan(&queryResults).Error

	if err != nil {
		logger.ApiLog().Error(err)
		return nil, &apihelpers.ApiError{
			Code:  fiber.StatusInternalServerError,
			Error: errors.New("internal server error"),
		}
	}

	for _, source := range queryResults {
		result = append(result, map[string]interface{}{
			"source":           source.Source,
			"sourceId":         source.SourceId,
			"verificationTime": source.VerificationTime.Unix(),
		})
	}
	return
}
