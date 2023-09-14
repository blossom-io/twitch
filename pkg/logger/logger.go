package logger

import (
	"context"
	"log/slog"
	"os"
)

type Logger interface {
	DebugContext(ctx context.Context, msg string, args ...any)
	Debug(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	Info(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	Warn(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
	Error(msg string, args ...any)
}

type logger struct {
	log *slog.Logger
}

var _ Logger = (*logger)(nil)

func New() Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	logger := &logger{
		log: log,
	}

	return logger
}

func (l *logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log.DebugContext(ctx, msg, args...)
}

func (l *logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}

func (l *logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log.InfoContext(ctx, msg, args...)
}

func (l *logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log.WarnContext(ctx, msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

func (l *logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log.ErrorContext(ctx, msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}
