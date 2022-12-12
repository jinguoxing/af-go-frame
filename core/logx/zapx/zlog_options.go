package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DefaultLogName = "zap"
	//ConsoleFormat log output format, console and json format
	ConsoleFormat = "console"
	JsonFormat    = "json"
)

const (
	ZapCallerSkip     = 1
	ProjectCallerSkip = 2
)

//Options  single logger config
type Options struct {
	Name            string `json:"name"               mapstructure:"name"`
	local           bool
	Default         bool         `json:"default"               mapstructure:"default"`
	EnableCaller    bool         `json:"enable_caller"      mapstructure:"enable_caller"`
	StacktraceLevel string       `json:"stacktrace_level"   mapstructure:"stacktrace_level"`
	Development     bool         `json:"development"        mapstructure:"development"`
	StackFilter     []string     `json:"stack_filter"  		mapstructure:"stack_filter"`
	CoreConfigs     []CoreConfig `json:"cores"       mapstructure:"cores"`
}

//DefaultOptions default config, only console core
func DefaultOptions() Options {
	consoleConfig := DefaultCoreConfig()
	return Options{
		Name:            DefaultLogName,
		local:           true,
		Default:         true,
		EnableCaller:    true,
		Development:     true,
		StacktraceLevel: zapcore.ErrorLevel.String(),
		CoreConfigs:     []CoreConfig{consoleConfig},
	}
}

//core logger output stream, config output file
func (opts Options) core() zapcore.Core {
	cores := make([]zapcore.Core, 0)
	for _, cfg := range opts.CoreConfigs {
		cores = append(cores, cfg.core())
	}
	//set default core
	if len(cores) <= 0 {
		cores = append(cores, DefaultCoreConfig().core())
	}
	return zapcore.NewTee(cores...)
}

//options gen slice of zap.Option
func (opts Options) options() []zap.Option {
	options := make([]zap.Option, 0)
	if opts.Development {
		options = append(options, zap.Development())
	}
	if opts.StacktraceLevel != "" {
		options = append(options, zap.AddStacktrace(genLogLevel(opts.StacktraceLevel)))
	}
	if opts.EnableCaller {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(ZapCallerSkip))
	}
	return options
}

func (opts Options) ZapLogger() *zapLogger {
	core := opts.core()
	options := opts.options()

	l := zap.New(core, options...)
	zap.RedirectStdLog(l)
	return &zapLogger{
		Config:    opts,
		zapLogger: l.Named(opts.Name),
	}
}
