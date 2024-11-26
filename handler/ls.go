package handler

import (
	"fmt"
	"os"
	"regexp"
	"simple-ls/logging"
)

// This function receives the list of command line arguments including `ls` and
// then processes accordingly
func HandleLs(arguments []string) {
    logging.InfoLog.Printf("Received arguments %s", arguments)
    validateArgs(arguments)
    if len(arguments) == 1 {
        handleSimpleLs()
    } else {
        handleLsWithFlags(arguments[1:])
    }
}

// Command line arguments should only contain letters and dashes
func validateArgs(args []string) {
    alphaNumericRegex := regexp.MustCompile(`^[a-zA-z-]+$`)
    for _, element := range args {
        if !alphaNumericRegex.MatchString(element) {
            logging.ErrorLog.Printf("Program argument %s contains non-alphabet character", element)
            panic("Invalid arguments found!")
        }
    }
}

// The simple ls functionality should be to list all files and directories in current directory
func handleSimpleLs() {
    logging.InfoLog.Println("Handling simple ls")
    currentWorkingDirectory, err := os.Getwd()
    if err != nil {
        logging.ErrorLog.Printf("Error encountered while getting current working directory: %s", err)
    }
    logging.InfoLog.Printf("Current working directory: %s", currentWorkingDirectory)
    files, err := os.Open(currentWorkingDirectory)
    if err != nil {
        logging.ErrorLog.Printf("Error encountered when opening directory %s", currentWorkingDirectory)
    }
    fileInfos, err := files.Readdir(-1)
    if err != nil {
        logging.ErrorLog.Printf("Error encountered while reading contents of %s directory: %s", currentWorkingDirectory, err)
    }
    formatOutput(fileInfos)
    logging.InfoLog.Println("Printed contents of directory on screen!")
}

func handleLsWithFlags(flags []string) {
    logging.InfoLog.Printf("Handling ls with flags: %s", flags)
}

func formatOutput(fileInfos []os.FileInfo) {
    var maxLen int
    for _, file := range fileInfos {
        maxLen = max(maxLen, len(file.Name()))
    }
    columnWidth := maxLen + 10

    fmt.Printf("%-*s%-*s%-*s%-*s\n", columnWidth, "Name", columnWidth, "Is Directory", columnWidth, "Size in Bytes", columnWidth, "Modification Time")
    for _, file := range fileInfos {
        fmt.Printf("%-*s%-*t%-*d%-*s\n", columnWidth, file.Name(), columnWidth, file.IsDir(), columnWidth, file.Size(), columnWidth, file.ModTime())
    }

}
