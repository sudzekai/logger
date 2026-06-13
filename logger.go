package logger

import "log/slog"

type Logger struct {
	*slog.Logger
}

func New(handler *slog.JSONHandler) *Logger {
	return &Logger{Logger: slog.New(handler)}
}

func (logger *Logger) Debug(
	msg string,
	args ...any,
) {
	logger.Logger.Debug(msg, args...)
}

func (logger *Logger) Info(
	msg string,
	args ...any,
) {
	logger.Logger.Info(msg, args...)
}

func (logger *Logger) Warn(
	msg string,
	args ...any,
) {
	logger.Logger.Warn(msg, args...)
}

func (logger *Logger) Error(
	msg string,
	args ...any,
) {
	logger.Logger.Error(msg, args...)
}

func (logger *Logger) With(args ...any) *Logger {
	return &Logger{
		Logger: logger.Logger.With(args...),
	}
}
