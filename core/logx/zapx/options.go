package zapx

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)
const (
	DefaultFileNameFormat = "-%Y-%m-%d.log"
	DefaultLogName        = "zap"
)

const (
	ConsoleFormat = "console"
	JsonFormat    = "json"
)
const (
	ConsoleDestination = "stdout"
)
const (
	ZapxCallerSkip    = 1
	ProjectCallerSkip = 2
)

type Options struct {
	Name            string        `json:"name"               mapstructure:"name"`
	EnableCaller    bool          `json:"enable-caller"     mapstructure:"enable-caller"`
	CallerSkip      int           `json:"caller-skip"        mapstructure:"caller-skip"`
	StacktraceLevel zapcore.Level `json:"disable-stacktrace" mapstructure:"disable-stacktrace"`
	Development     bool          `json:"development"        mapstructure:"development"`
	CoreConfigs     []CoreConfig  `json:"core-configs"       mapstructure:"core-configs"`
}

type CoreConfig struct {
	RotateSize  int64                `json:"rotate_size" mapstructure:"rotate_size"`
	Destination string               `json:"destination" mapstructure:"destination"`
	Format      string               `json:"format" mapstructure:"format"`
	LogLevel    zap.LevelEnablerFunc `json:"log_level" mapstructure:"log_level"`
}

func DefaultInfoLevel() zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel && lvl <= zapcore.WarnLevel
	})
}

func DefaultErrorLevel() zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel && lvl <= zapcore.FatalLevel
	})
}

func DefaultConsoleLevel() zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
}

func (opts Options) NewZapLogger() *zapLogger {
	cores := make([]zapcore.Core, 0)
	for _, cfg := range opts.CoreConfigs {
		cores = append(cores, cfg.newCore())
	}
	core := zapcore.NewTee(cores...)
	options := make([]zap.Option, 0)

	if opts.Development {
		options = append(options, zap.Development())
	}
	if opts.StacktraceLevel >= 0 {
		options = append(options, zap.AddStacktrace(opts.StacktraceLevel))
	}
	if opts.EnableCaller {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(opts.CallerSkip))
	}

	l := zap.New(core, options...)
	logger := &zapLogger{
		zapLogger: l.Named(opts.Name),
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}
	zap.RedirectStdLog(l)
	return logger
}

func NewDefaultOptions() Options {
	infoConfig := CoreConfig{
		RotateSize:  10 * MB,
		Destination: "logs/info.log",
		Format:      JsonFormat,
		LogLevel:    DefaultInfoLevel(),
	}
	errorConfig := CoreConfig{
		RotateSize:  10 * MB,
		Destination: "logs/error.log",
		Format:      ConsoleFormat,
		LogLevel:    DefaultErrorLevel(),
	}
	consoleConfig := CoreConfig{
		RotateSize:  10 * MB,
		Destination: ConsoleDestination,
		Format:      ConsoleFormat,
		LogLevel:    DefaultConsoleLevel(),
	}
	return Options{
		Name:            DefaultLogName,
		EnableCaller:    true,
		CallerSkip:      ZapxCallerSkip,
		Development:     true,
		StacktraceLevel: PanicLevel,
		CoreConfigs:     []CoreConfig{infoConfig, errorConfig, consoleConfig},
	}
}

func (c CoreConfig) newEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// when output to local path, with color is forbidden
	if c.Format == ConsoleFormat {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func (c CoreConfig) newWriteSyncer() zapcore.WriteSyncer {
	switch c.Destination {
	case ConsoleDestination:
		return zapcore.Lock(os.Stdout)
	default:
		hook, _ := rotatelogs.New(
			strings.Replace(c.Destination, ".log", "", -1)+DefaultFileNameFormat,
			rotatelogs.WithLinkName(c.Destination),
			rotatelogs.WithRotationSize(c.RotateSize),
		)
		return zapcore.AddSync(hook)
	}
}

func (c CoreConfig) newCore() zapcore.Core {
	return zapcore.NewCore(c.newEncoder(), zapcore.AddSync(c.newWriteSyncer()), c.LogLevel)
}
