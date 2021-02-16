package logger

import (
	"path/filepath"
	"runtime"
)

func init() {
	//get current directory
	_, file, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(file)

	// server loggers
	initApiLog(currentPath)

	// logs HTTP request/response details.
	initRequestLog(currentPath)
}
