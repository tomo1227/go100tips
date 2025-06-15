package main

import (
	"log/slog"
	"os"
)

type Logger struct {
	slog *slog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		slog: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (l *Logger) Info(msg string, args ...any) {
	l.slog.Info(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.slog.Error(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.slog.Warn(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.slog.Debug(msg, args...)
}
