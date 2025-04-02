package services

import (
	"fmt"
	"log"
	"os"
)

type StdLogService struct {
	warn  *log.Logger
	error *log.Logger
	info  *log.Logger
}

func NewStdLogService() *StdLogService {
	return &StdLogService{
		warn:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime),
		error: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
		info:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
	}
}

func (l *StdLogService) Warnf(msg string, args ...interface{}) {
	l.warn.Println(fmt.Sprintf(msg, args...))
}

func (l *StdLogService) Errorf(msg string, args ...interface{}) {
	l.error.Println(fmt.Sprintf(msg, args...))
}

func (l *StdLogService) Infof(msg string, args ...interface{}) {
	l.info.Println(fmt.Sprintf(msg, args...))
}
