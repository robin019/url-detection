package service

import (
	"reflect"
	"time"

	"github.com/robin019/url-detection/src/psersistence/sql"

	"github.com/robin019/url-detection/src/apihelpers"
)

func UrlCheck(params interface{}) (result []map[string]interface{}, err *apihelpers.ApiError) {
	value := reflect.ValueOf(params).Elem()
	url := value.FieldByName("Url").String()

	// crop the last character if it is a '/'
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
	db.Raw(`select md.source, md.source_id, md.verification_time
			from malicious_url AS mu JOIN malicious_url_detail AS md on mu.id = md.url_id
			where md5(lower(mu.url)) = md5(lower(?))`, url).Scan(&queryResults)

	for _, source := range queryResults {
		result = append(result, map[string]interface{}{
			"source":           source.Source,
			"sourceId":         source.SourceId,
			"verificationTime": source.VerificationTime.Unix(),
		})
	}
	return
}
