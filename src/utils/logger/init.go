package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/robin019/url-detection/src/utils/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	apiLog *zap.SugaredLogger
)

func init() {
	//get current directory
	_, file, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(file)

	initApiLog(currentPath)
}

func ApiLog() *zap.SugaredLogger {
	return apiLog
}

func initApiLog(currentPath string) {
	//create .log file if not exists
	logPath := fmt.Sprintf(`%s/../../../log/%s`, currentPath, config.Log.File)
	err := touchFile(logPath)
	if err != nil {
		log.Fatal("can't initialize zap logger : " + err.Error())
	}

	logBuilder := zap.NewDevelopmentConfig()
	logBuilder.EncoderConfig.EncodeTime = syslogTimeEncoder
	logBuilder.OutputPaths = []string{
		logPath,
	}

	logger, err := logBuilder.Build()
	if err != nil {
		log.Fatal("can't initialize zap logger : " + err.Error())
	}
	apiLog = logger.Sugar()
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

//create log file if not exists
func touchFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}
