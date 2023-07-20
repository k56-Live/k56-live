package common

import (
	"fmt"
	"log"
	"os"
)

// Logger is a custom logger with different log levels
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLogger creates a new instance of the custom logger
func NewLogger() *Logger {
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}
}

// Info logs information messages
func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

// Error logs error messages
func (l *Logger) Error(err error) {
	if err != nil {
		l.errorLogger.Println(err)
	}
}
