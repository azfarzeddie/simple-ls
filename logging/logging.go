package logging

import (
	"log"
	"os"
)

var (
    InfoLog *log.Logger
    WarningLog *log.Logger
    ErrorLog *log.Logger
)

func InitLogs() {
    logFileName := "logs/application-logs.log"
    logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    InfoLog = log.New(logFile, "INFO: " , log.LstdFlags)
    WarningLog = log.New(logFile, "WARNING: " , log.LstdFlags)
    ErrorLog = log.New(logFile, "ERROR: " , log.LstdFlags)
}
