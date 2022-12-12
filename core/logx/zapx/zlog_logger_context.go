package zapx

import (
	"context"
	"go.uber.org/zap"
)

type key int

const (
	logContextKey key = iota
)

// SetContext method output with specified context value.
func SetContext(ctx context.Context) Logger {
	return std.SetContext(ctx)
}

//SetContext add 'requestID','userName','watcher' Field in log
func (l *zapLogger) SetContext(ctx context.Context) Logger {
	lg := l.clone()

	if requestID := ctx.Value(KeyRequestID); requestID != nil {
		lg.zapLogger = lg.zapLogger.With(zap.Any(KeyRequestID, requestID))
	}
	if username := ctx.Value(KeyUsername); username != nil {
		lg.zapLogger = lg.zapLogger.With(zap.Any(KeyUsername, username))
	}
	if watcherName := ctx.Value(KeyWatcherName); watcherName != nil {
		lg.zapLogger = lg.zapLogger.With(zap.Any(KeyWatcherName, watcherName))
	}

	return lg
}
