package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
	prefix string
}

// Create a new logger instance with service name as prefix
func New(service string) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "["+service+"] ", log.Ldate|log.Ltime|log.Lmicroseconds),
		prefix: service,
	}
}

// Info logs an info message
func (l *Logger) Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.Printf("INFO: "+msg, args...)
	} else {
		l.Println("INFO: " + msg)
	}
}

// Error logs an error message
func (l *Logger) Error(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.Printf("ERROR: "+msg, args...)
	} else {
		l.Println("ERROR: " + msg)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.Printf("DEBUG: "+msg, args...)
	} else {
		l.Println("DEBUG: " + msg)
	}
}
