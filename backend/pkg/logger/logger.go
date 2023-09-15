package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

func Init() {
	runLogFile, _ := os.OpenFile(
		getLogFilename(),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
}

func getLogFilename() string {
	now := time.Now()
	formattedTime := now.Format("02-01-06-15-04-05")
	filename := fmt.Sprintf("./logs/logs_%s", formattedTime)
	return filename
}
