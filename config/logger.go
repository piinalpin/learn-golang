package config

import (
	"io"
	"os"
	"path"
	"runtime"
	"strconv"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
)

func InitLog() {
	godotenv.Load()
	filename := os.Getenv("logger.output-filename")
	if len(filename) > 0 {
		f, _ := os.OpenFile("logs/" + filename, os.O_CREATE|os.O_WRONLY, 0777)
		log.SetOutput(io.MultiWriter(os.Stderr, f))
	}

	log.SetLevel(getLoggerLevel(os.Getenv("logger.log-level")))
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
		CustomCallerFormatter: func(frame *runtime.Frame) (file string) {
			fileName := " [" + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line) + "]"
			return fileName
		},
		CallerFirst: true,
	})

}

func getLoggerLevel(value string) log.Level {
	switch value {
	case "DEBUG":
		return log.DebugLevel
	case "TRACE":
		return log.TraceLevel
	default:
		return log.InfoLevel
	}
}
