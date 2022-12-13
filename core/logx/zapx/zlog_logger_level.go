package zapx

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level is an alias for the level structure in the underlying log frame.
type Level = zapcore.Level

var (
	DisableLevel = zapcore.DebugLevel - 1
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

const (
	DefaultLogLevelString = "info"
	DefaultLogLevel       = zapcore.InfoLevel
)

var logLevelMap = map[string]zapcore.Level{
	zapcore.DebugLevel.String(): zapcore.DebugLevel,
	zapcore.InfoLevel.String():  zapcore.InfoLevel,
	zapcore.WarnLevel.String():  zapcore.WarnLevel,
	zapcore.ErrorLevel.String(): zapcore.ErrorLevel,
	zapcore.PanicLevel.String(): zapcore.PanicLevel,
	zapcore.FatalLevel.String(): zapcore.FatalLevel,
}

//genLogLevel gen log level zapcore.Level, Will return 'zapcore.InfoLevel' if not found
func genLogLevel(l string) zapcore.Level {
	level, ok := logLevelMap[l]
	if !ok {
		return DefaultLogLevel
	}
	return level
}

//genLogLevelFunc gen 'zap.LevelEnablerFunc' for Options
func genLogLevelFunc(l string) zap.LevelEnablerFunc {
	level := genLogLevel(l)
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})
}

//SetLevel  set core log level,
func (l *zapLogger) SetLevel(level Level, destinations ...string) Logger {
	dest := ""
	if len(destinations) > 0 && destinations[0] != "" {
		dest = destinations[0]
	}
	exists := false
	for i, cc := range l.Config.CoreConfigs {
		if cc.Destination == dest && dest != "" {
			exists = true
		}
		// find the core
		if cc.Destination == dest && dest != "" {
			l.Config.CoreConfigs[i].Destination = dest
			l.Config.CoreConfigs[i].LogLevel = level.String()
		}
	}
	if !exists && dest != "" {
		panic(fmt.Errorf("destination %s not exists", dest))
	}
	return l.Reload()
}

//GetLevel  set core log level,
func (l *zapLogger) GetLevel(destinations ...string) Level {
	dest := ""
	if len(destinations) > 0 && destinations[0] != "" {
		dest = destinations[0]
	}
	if len(l.Config.CoreConfigs) == 1 {
		return genLogLevel(l.Config.CoreConfigs[0].LogLevel)
	}
	level := DisableLevel
	for _, cc := range l.Config.CoreConfigs {
		if cc.Destination == dest {
			level = genLogLevel(cc.LogLevel)
		}
	}

	return level
}

// CheckIntLevel used for other log wrapper such as klog which return if logging a
// message at the specified level is enabled.
func CheckIntLevel(level int32) bool {
	var lvl zapcore.Level
	if level < 5 {
		lvl = zapcore.InfoLevel
	} else {
		lvl = zapcore.DebugLevel
	}
	checkEntry := std.zapLogger.Check(lvl, "")

	return checkEntry != nil
}
