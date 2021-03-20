package main

import (
	"fmt"
	"os"
)

//SeverityType - A type to hold error severity.
//Critical will cause immediate program exit
//Warning will print a warning to screen but will not exit
//Info will only print a warning to log if specified
type SeverityType string

const (
	Critical SeverityType = "Critical" //Critical - print error message to stdout and exit immediately
	Warning  SeverityType = "Warning" //Warning - Print warning message to stdout but do not exit
	Info     SeverityType = "Info" //Info - Print info message to log (if provided)
)

//CheckStats - A method to handle return status
func CheckStatus(ok bool, Message string, Severity SeverityType) {
	if !ok {
		handleError(fmt.Sprintf("%s: %s", Severity, Message), Severity)
	}
}


//CheckError - A method to handle errors
func CheckError(Error error, Message string, Severity SeverityType) {
	if Error != nil {
		handleError(fmt.Sprintf("%s: %s - %s", Severity, Message, Error), Severity)
	}
}

func handleError(Message string, Severity SeverityType) {
	fmt.Println()
	if Severity == Critical {
		os.Exit(1)
	}
}
