package adapters

import (
	"fmt"
	"log/slog"
	"os"
)

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() *SlogLogger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)

	return &SlogLogger{logger: logger}
}

func (l *SlogLogger) Warnf(msg string, args ...interface{}) {
	l.logger.Warn(format(msg, args...))
}

func (l *SlogLogger) Errorf(msg string, args ...interface{}) {
	l.logger.Error(format(msg, args...))
}

func (l *SlogLogger) Infof(msg string, args ...interface{}) {
	l.logger.Info(format(msg, args...))
}

func format(msg string, args ...interface{}) string {
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}
