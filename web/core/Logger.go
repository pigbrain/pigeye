package core

import (
	"io"
	"log"
)

var logger *log.Logger

func Create(writer io.Writer) {
	logger = log.New(writer, "EYE: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetLogger() *log.Logger {
	return logger
}
