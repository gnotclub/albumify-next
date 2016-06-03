package util

import (
	"log"
	"os"
)

var Logger *log.Logger

func GetLogger() {
	Logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
}
