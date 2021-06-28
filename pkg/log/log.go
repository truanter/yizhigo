package log

import (
	"github.com/sirupsen/logrus"
	"github.com/truanter/yizhigo/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var Logger = logrus.New()
var output io.Writer

func initOutput(maxSizeMB, maxBackups, maxDays int) {
	if !config.IsProdEnv() {
		output = io.MultiWriter(os.Stdout)
		return
	}
	filePath := config.GetLogFilePath()
	output = io.MultiWriter(
		&lumberjack.Logger{
			Filename:   filePath,
			MaxSize:    maxSizeMB,
			MaxAge:     maxDays,
			MaxBackups: maxBackups,
			LocalTime:  true,
			Compress:   false,
		},
		os.Stdout,
	)
}

func init() {
	initOutput(config.MaxLogFileSize, config.MaxLogFileBackup, config.MaxLogFileDays)
	Logger.SetFormatter(
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
	if config.IsProdEnv() {
		Logger.SetLevel(logrus.InfoLevel)
	} else {
		Logger.SetLevel(logrus.DebugLevel)
	}
	Logger.SetOutput(output)
}
