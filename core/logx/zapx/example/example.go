package main

import (
	"context"
	"github.com/jinguoxing/af-go-frame/core/logx/zapx"
)

var (
	h bool

	level  int
	format string
)

func main() {
	// 初始化全局logger
	defer zapx.Flush()

	// Debug、Info(with field)、Warnf、Errorw使用
	zapx.Debug("This is a debug message")
	zapx.Info("This is a info message", zapx.Int32("int_key", 10))
	zapx.Warnf("This is a formatted %s message", "warn")
	zapx.Errorw("Message printed with Errorw", "X-Request-ID", "fbf54504-64da-4088-9b86-67824a7fb508")

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
	zapx.V(zapx.InfoLevel).Info("This is a V level message")
	zapx.V(zapx.ErrorLevel).
		Infow("This is a V level message with fields", "X-Request-ID", "7a7b9f24-4cae-4b2a-9464-69088b45b904")
}
