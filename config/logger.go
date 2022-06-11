package config

import (
	"io"
	"os"
	"path"
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"
	nested "github.com/antonfisher/nested-logrus-formatter"
)

func InitLog() {
	f, _ := os.OpenFile("logs/learn-rest-api.log", os.O_CREATE|os.O_WRONLY, 0777)
	log.SetOutput(io.MultiWriter(os.Stderr, f))

	log.SetLevel(log.TraceLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel: true,
		CustomCallerFormatter: func(frame *runtime.Frame) (file string) {
			fileName := " [" + path.Base(frame.File) + ":" + strconv.Itoa(frame.Line) + "]"
			return fileName
		},
		CallerFirst: true,
	})
	
}