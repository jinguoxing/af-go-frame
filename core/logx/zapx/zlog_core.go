package zapx

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
	"strings"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)
const (
	KBFormat = "KB"
	MBFormat = "MB"
	GBFormat = "GB"
)

const (
	//DefaultFileNameFormat name format for file rotate
	DefaultFileNameFormat = "-%Y-%m-%d.log"
	//DefaultRotateSize default file rotate size
	DefaultRotateSize = "10M"
)

const (
	ConsoleDestination = "stdout"
)

const (
	StdoutCore    = "stdout"
	LocalFileCore = "file"
)

//CoreConfig zap log core config
type CoreConfig struct {
	RotateSize     string `json:"rotate_size" mapstructure:"rotate_size"`         //file rotate size
	Destination    string `json:"destination" mapstructure:"destination"`         //log output destination
	CoreType       string `json:"core_type" mapstructure:"core_type"`             //core writeSyncer type, 'file', 'stdout'
	EnableColor    bool   `json:"enable_color" mapstructure:"enable_color"`       //enable color
	OutputFormat   string `json:"output_format" mapstructure:"output_format"`     //log line output format: 'line', 'json'
	FileNameFormat string `json:"filename_format" mapstructure:"filename_format"` //name format in file rotate
	LogLevel       string `json:"log_level" mapstructure:"log_level"`             //lowest log level in this core
}

//DefaultCoreConfig default console config
func DefaultCoreConfig() CoreConfig {
	return CoreConfig{
		CoreType:     StdoutCore,
		RotateSize:   DefaultRotateSize,
		EnableColor:  true,
		Destination:  ConsoleDestination,
		OutputFormat: ConsoleFormat,
		LogLevel:     DefaultLogLevelString,
	}
}

//core generate 'zapcore.Core'
func (c CoreConfig) core() zapcore.Core {
	return zapcore.NewCore(c.encoder(), c.writeSyncer(), genLogLevelFunc(c.LogLevel))
}

//encoder generate 'zapcore.Encoder'
func (c CoreConfig) encoder() zapcore.Encoder {
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
	if c.Destination == ConsoleDestination && c.EnableColor {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	if c.OutputFormat == ConsoleFormat {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

//writeSyncer query 'zapcore.WriteSyncer' by coreType
func (c CoreConfig) writeSyncer() zapcore.WriteSyncer {
	switch c.Destination {
	case StdoutCore:
		return c.StdoutWriteSyncer()
	case LocalFileCore:
		return c.FileWriteSyncer()
	default:
		return c.GetWriteSyncer(c.CoreType)
	}
}

//parseSize parse strings size format like '10KB', '10MB', '10GB'.
//also support lower case like '10Kb', '10mb', '10gB'
func parseSize(sizeFormat string) (s int64) {
	size := strings.TrimSpace(sizeFormat)
	if size == "" {
		return 0
	}
	size = strings.ToUpper(size)
	var err error
	var format int64
	switch {
	case strings.Contains(size, KBFormat):
		s, err = strconv.ParseInt(strings.TrimSuffix(size, KBFormat), 10, 64)
		format = KB
	case strings.Contains(size, MBFormat):
		s, err = strconv.ParseInt(strings.TrimSuffix(size, MBFormat), 10, 64)
		format = MB
	case strings.Contains(size, GBFormat):
		s, err = strconv.ParseInt(strings.TrimSuffix(size, GBFormat), 10, 64)
		format = GB
	default:
		s, err = strconv.ParseInt(size, 10, 64)
	}
	if err != nil {
		panic(fmt.Sprintf("invalid size format '%s'", sizeFormat))
	}
	return s * format
}

//checkDir check weather 'dstDir' exists or not, will panic if even this 'dstDir' can not be built
func checkDir(dstDir string) {
	if dstDir == "" {
		panic("empty config")
	}
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		if err = os.MkdirAll(dstDir, os.ModePerm); err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("create dir: %s", dstDir)
		}
	}
}
