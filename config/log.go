package config

import (
	"fmt"
	"os"
)

const (
	MaxLogFileSize   = 200
	MaxLogFileBackup = 512
	MaxLogFileDays   = 7
	logFilePath      = "/yizhigo/log/inst-%s.log"
)

func GetLogFilePath() string {
	instNO := os.Getenv("inst_no")
	if instNO == "" {
		instNO = "1"
	}
	return fmt.Sprintf(logFilePath, instNO)
}
