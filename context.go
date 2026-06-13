package logger

import "context"

type contextKey string

var LoggerKey contextKey = "logger"

func SetContextKey(key contextKey) {
	LoggerKey = key
}

func FromContext(ctx context.Context) *Logger {
	val := ctx.Value(LoggerKey)

	log, ok := val.(*Logger)

	if ok {
		return log
	} else {
		return nil
	}
}

func ToContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, logger)
}
