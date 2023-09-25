package contextutils

import (
	"context"

	"go.uber.org/zap"
)

const (
	loggerKey = "logger"
)

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func Logger(ctx context.Context) *zap.SugaredLogger {
	v := ctx.Value(loggerKey)
	if v == nil {
		return zap.NewNop().Sugar()
	}

	return v.(*zap.SugaredLogger)
}
