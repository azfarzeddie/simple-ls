package main

import (
	"os"
	"simple-ls/handler"
	"simple-ls/logging"
	"strings"
)

func main() {
	logging.InitLogs()
	logging.InfoLog.Println("Initiated logging. Proceeding to starting application.")
	programArgs := os.Args
	logging.InfoLog.Printf("Command executed: %s", strings.Join(programArgs, " "))
    logging.InfoLog.Printf("Arguments: %q", programArgs)
    handler.HandleLs(programArgs[1:])
}
