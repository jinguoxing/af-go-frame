package zapx

import (
	"context"
	"go.uber.org/zap"
)

// zapLogger is a logr.Logger that uses Zap to log.
type zapLogger struct {
	// NB: this looks very similar to zap.SugaredLogger, but
	// deals with our desire to have multiple verbosity levels.
	zapLogger *zap.Logger
	Config    Options
}

// GetLogger get logger by name, will return a default logger if not exists
func GetLogger(names ...string) *zapLogger {
	name := ""
	if len(names) > 0 && names[0] != "" {
		name = names[0]
	}
	log, ok := loggerMap[name]
	if ok {
		return log
	}
	std.Warnf("logger '%s' not exists, return the default logger", name)
	return std
}

// SugaredLogger returns global sugared logger.
func SugaredLogger() *zap.SugaredLogger {
	return std.zapLogger.Sugar()
}

//nolint:predeclared
func (l *zapLogger) clone() *zapLogger {
	copy := *l
	return &copy
}

func (l *zapLogger) WithName(name string) Logger {
	newLogger := l.zapLogger.Named(name)
	return &zapLogger{
		Config:    l.Config,
		zapLogger: newLogger,
	}
}

func (l *zapLogger) Flush() {
	_ = l.zapLogger.Sync()
}

func (l *zapLogger) Write(p []byte) (n int, err error) {
	l.zapLogger.Info(string(p))

	return len(p), nil
}

func (l *zapLogger) WithValues(keysAndValues ...interface{}) Logger {
	newLogger := l.zapLogger.With(handleFields(l.zapLogger, keysAndValues)...)
	return &zapLogger{
		Config:    l.Config,
		zapLogger: newLogger,
	}
}

func (l *zapLogger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, logContextKey, l)
}

func (l *zapLogger) Debug(msg string, fields ...Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *zapLogger) Debugf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Debugf(format, v...)
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Info(msg string, fields ...Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *zapLogger) Infof(format string, v ...interface{}) {
	l.zapLogger.Sugar().Infof(format, v...)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

func (l *zapLogger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *zapLogger) Warnf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Warnf(format, v...)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Error(msg string, fields ...Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *zapLogger) Errorf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Errorf(format, v...)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Panic(msg string, fields ...Field) {
	l.zapLogger.Panic(msg, fields...)
}

func (l *zapLogger) Panicf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Panicf(format, v...)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

func (l *zapLogger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatal(msg, fields...)
}

func (l *zapLogger) Fatalf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Fatalf(format, v...)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

// PrintStack prints the caller stack,
// the optional parameter `skip` specify the skipped stack offset from the end point.
func (l *zapLogger) PrintStack(skip ...int) {
	s := l.GetStack(skip...)
	l.Info("Stack:\n" + s)
}

// GetStack returns the caller stack content,
// the optional parameter `skip` specify the skipped stack offset from the end point.
func (l *zapLogger) GetStack(skip ...int) string {
	stackSkip := ZapCallerSkip
	if len(skip) > 0 {
		stackSkip += skip[0]
	}
	filters := []string{pathFilterKey}
	if len(l.Config.StackFilter) > 0 {
		filters = append(filters, l.Config.StackFilter...)
	}
	return StackWithFilters(filters, stackSkip)
}

// handleFields converts a bunch of arbitrary key-value pairs into Zap fields.  It takes
// additional pre-converted Zap fields, for use with automatically attached fields, like
// `error`.
func handleFields(l *zap.Logger, args []interface{}, additional ...zap.Field) []zap.Field {
	// a slightly modified version of zap.SugaredLogger.sweetenFields
	if len(args) == 0 {
		// fast-return if we have no suggared fields.
		return additional
	}

	// unlike Zap, we can be pretty sure users aren't passing structured
	// fields (since logr has no concept of that), so guess that we need a
	// little less space.
	fields := make([]zap.Field, 0, len(args)/2+len(additional))
	for i := 0; i < len(args); {
		// check just in case for strongly-typed Zap fields, which is illegal (since
		// it breaks implementation agnosticism), so we can give a better error message.
		if _, ok := args[i].(zap.Field); ok {
			l.DPanic("strongly-typed Zap Field passed to logr", zap.Any("zap field", args[i]))

			break
		}

		// make sure this isn't a mismatched key
		if i == len(args)-1 {
			l.DPanic("odd number of arguments passed as key-value pairs for logging", zap.Any("ignored key", args[i]))

			break
		}

		// process a key-value pair,
		// ensuring that the key is a string
		key, val := args[i], args[i+1]
		keyStr, isString := key.(string)
		if !isString {
			// if the key isn't a string, DPanic and stop logging
			l.DPanic(
				"non-string key argument passed to logging, ignoring all later arguments",
				zap.Any("invalid key", key),
			)

			break
		}

		fields = append(fields, zap.Any(keyStr, val))
		i += 2
	}

	return append(fields, additional...)
}
