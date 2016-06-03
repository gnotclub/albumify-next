package util

import (
	"log"
	"os"
)

var Logger *log.Logger

// Returns a new logger instance
func GetLogger() {
	Logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
}
