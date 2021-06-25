package vbmcgoalie

import (
	"log"
)

// logStdOut just logs something to stdout
func logStdOut(s string) {
	log.Printf("%s\n", string(s))
}

// logStdErr just logs to stderr
func logStdErr(s string) {
	log.Fatalf("%s\n", string(s))
}

// Stoerr wraps a string in an error object
func Stoerr(s string) error {
	return &errorString{s}
}

// logStringAsErrorToStdout wraps a string in an error object and pushes to stdout
func logStringAsErrorToStdout(s string) {
	check(&errorString{s})
}

// logStringAsErrorToStderr wraps a string in an error object and pushes to stdout
func logStringAsErrorToStderr(s string) {
	checkAndFail(&errorString{s})
}

// check does error checking
func check(e error) {
	if e != nil {
		log.Printf("error: %v", e)
	}
}

// checkAndFail checks for an error type and fails
func checkAndFail(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}
