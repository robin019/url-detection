package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/robin019/url-detection/src/utils/config"
)

var (
	RequestLog *os.File
)

func initRequestLog(currentPath string) {
	logPath := fmt.Sprintf(`%s/../../../log/%s`, currentPath, config.Log.RequestLogFile)
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	RequestLog = file
}
