package main

import (
	"context"
	"fmt"
	"github.com/jinguoxing/af-go-frame/core/config"
	"github.com/jinguoxing/af-go-frame/core/config/sources/env"
	"github.com/jinguoxing/af-go-frame/core/config/sources/file"
	"github.com/jinguoxing/af-go-frame/core/logx/zapx"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"runtime"
	"strings"
)

var pwd string = "."

func init() {
	os.Setenv(config.ProjectEnvKey, "DEV")

	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
		pwd = abPath
	}
}

//InitSources init config source, with env
func InitSources(paths ...string) {
	sources := make([]config.Source, len(paths)+1)
	sources[0] = env.NewSource()
	for i, path := range paths {
		sources[i+1] = file.NewSource(path)
	}
	config.Init(sources...)
}

//WriteSyncer 自定义的WriteSyncer, 如果需要其他的参数,可以在配置中增加
func WriteSyncer() zapcore.WriteSyncer {
	// check file path, create file if not exists
	destination := pwd + "/error.log"
	// set filename format
	fileNameFormat := zapx.DefaultFileNameFormat

	hook, _ := rotatelogs.New(
		strings.Replace(destination, ".log", "", -1)+fileNameFormat,
		rotatelogs.WithLinkName(destination),
		rotatelogs.WithRotationSize(10*zapx.MB),
	)
	return zapcore.AddSync(hook)
}

func main() {
	zapx.RegisterWriteSyncer("test", WriteSyncer())
	InitSources(fmt.Sprintf("%s/config.yaml", pwd))
	// 初始化全局logger
	defer zapx.Flush()

	logs := config.Scan[zapx.LogConfigs]()
	zapx.Loads(logs)

	// Debug、Info(with field)、Warnf、Errorw使用
	zapx.Debug("This is a debug message")
	zapx.Info("This is a info message")
	zapx.Warn("This is a formatted %s message")
	zapx.Error("Message printed with Errorw")
	//zapx.Panic("This is a panic message")

	zapx.Debugf("This is a %s message", "debug")
	zapx.Infof("This is a %s message", "info")
	zapx.Warn("This is a formatted %s message", zapx.Int32("int_key", 10))
	zapx.Errorw("Message printed with Errorw", "X-Request-ID", "fbf54504-64da-4088-9b86-67824a7fb508")
	zapx.Panic("This is a panic message")

	//
	zapx.Debug("This is a debug message")
	zapx.Info("This is a info message", zapx.Int32("int_key", 10))
	zapx.Warnf("This is a formatted %s message", "warn")
	zapx.Errorw("Message printed with Errorw", "X-Request-ID", "fbf54504-64da-4088-9b86-67824a7fb508")
	zapx.Panicf("This is a panic message")

	// WithValues使用
	lv := zapx.WithValues("X-Request-ID", "7a7b9f24-4cae-4b2a-9464-69088b45b904")
	lv.Infow("Info message printed with [WithValues] logger")
	lv.Infow("Debug message printed with [WithValues] logger")

	// Context使用

	ctx := lv.WithContext(context.Background())
	lc := zapx.FromContext(ctx)
	lc.Info("Message printed with [WithContext] logger")

	ln := lv.WithName("test")
	ln.Info("Message printed with [WithName] logger")

	// V level使用
	context.WithValue(ctx, "requestID", "RID:123456789156487")
	ln.SetLevel(zapx.InfoLevel).Info("This is a V level message")
	ln.Infow("log level %s", ln.GetLevel().String())
	ln.SetContext(ctx).Errorw("This is a V level message with fields")
	ln.Info("xxxxxxxxxxxx")
	ln.PrintStack(0)
	zapx.PrintStack()

	ln.Caller(true)
	ln.Development()

	zapx.GetLogger().RegisterHook(zapx.SampleHook)
	zapx.Info("has hook")

}
