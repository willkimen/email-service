package sloglogger

import (
	"log/slog"
)

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(logger *slog.Logger) *SlogLogger {
	return &SlogLogger{logger: logger}
}

func (l *SlogLogger) Info(msg string, fields ...any) {
	l.logger.Info(msg, fields...)
}

func (l *SlogLogger) Error(msg string, err error, fields ...any) {
	allFields := append(fields, "error", err)
	l.logger.Error(msg, allFields...)
}
