package zapx

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

var (
	std       = DefaultOptions().ZapLogger()
	loggerMap = make(map[string]*zapLogger)
	mu        sync.Mutex
)

//Logger check interface implementation
var _ Logger = &zapLogger{}

type Configurable interface {
	//Reload config return a new logger
	Reload() Logger
	//SetLevel  set core log level,
	SetLevel(level Level, destinations ...string) Logger
	GetLevel(destinations ...string) Level

	//Caller  set Caller
	Caller(flag bool)
	Development()

	//RegisterHook set hook
	RegisterHook(hooks ...func(zapcore.Entry) error)
	WithHook(hooks ...func(zapcore.Entry) error) Logger
}

// Logger represents the ability to log messages, both errors and not.
type Logger interface {
	Debug(msg string, fields ...Field)
	Debugf(format string, v ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Info(msg string, fields ...Field)
	Infof(format string, v ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warn(msg string, fields ...Field)
	Warnf(format string, v ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Error(msg string, fields ...Field)
	Errorf(format string, v ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panic(msg string, fields ...Field)
	Panicf(format string, v ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatal(msg string, fields ...Field)
	Fatalf(format string, v ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	//Configurable config method
	Configurable

	//PrintStack print stack
	PrintStack(skip ...int)
	GetStack(skip ...int) string

	Write(p []byte) (n int, err error)

	// WithValues adds some key-value pairs of context to a logger.
	// See Info for documentation on how key/value pairs work.
	WithValues(keysAndValues ...interface{}) Logger

	// WithName adds a new element to the logger's name.
	// Successive calls with WithName continue to append
	// suffixes to the logger's name.  It's strongly recommended
	// that name segments contain only letters, digits, and hyphens
	// (see the package documentation for more information).
	WithName(name string) Logger

	// WithContext returns a copy of context in which the log value is set.
	WithContext(ctx context.Context) context.Context
	SetContext(ctx context.Context) Logger

	// Flush calls the underlying Core's Sync method, flushing any buffered
	// log entries. Applications should take care to call Sync before exiting.
	Flush()
}

// StdErrLogger returns logger of standard library which writes to supplied zap
// logger at error level.
func StdErrLogger() *log.Logger {
	if std == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(std.zapLogger, zapcore.ErrorLevel); err == nil {
		return l
	}

	return nil
}

// StdInfoLogger returns logger of standard library which writes to supplied zap
// logger at info level.
func StdInfoLogger() *log.Logger {
	if std == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(std.zapLogger, zapcore.InfoLevel); err == nil {
		return l
	}

	return nil
}

// ZapLogger used for other log wrapper such as klog.
func ZapLogger() *zap.Logger {
	return std.zapLogger
}

func DefaultLogger() Logger {
	return std
}

// FromContext returns the value of the log key on the ctx.
func FromContext(ctx context.Context) Logger {
	if ctx != nil {
		logger := ctx.Value(logContextKey)
		if logger != nil {
			return logger.(Logger)
		}
	}

	return WithName("Unknown-Context")
}

//LogConfigs  config struct
type LogConfigs struct {
	Logs []Options `json:"logs"`
}

//Load create logger by config
func Load(cs ...Options) {
	hasDefault := false
	for _, c := range cs {
		if c.Name == "" {
			panic("missing logger name")
		}
		_, ok := loggerMap[c.Name]
		if ok {
			panic(fmt.Errorf("dulicated logger name '%s'", c.Name))
		}
		if c.Default && !std.Config.local {
			panic(fmt.Errorf("duplicated default logger '%s'", c.Name))
		}
		loggerMap[c.Name] = c.ZapLogger()
		if c.Default {
			hasDefault = true
			std = loggerMap[c.Name]
		}
	}
	if len(loggerMap) <= 0 {
		loggerMap[std.Config.Name] = std
	}
	if !hasDefault && len(cs) > 0 {
		std = loggerMap[cs[0].Name]
	}
}

//Loads create loggers by config
func Loads(cs LogConfigs) {
	Load(cs.Logs...)
}

//Reload config
func (l *zapLogger) Reload() Logger {
	newLogger := l.Config.ZapLogger()
	_, ok := loggerMap[l.Config.Name]
	if ok {
		l.Flush()
		l.zapLogger = newLogger.zapLogger
	} else {
		loggerMap[l.Config.Name] = newLogger
	}
	return l
}
